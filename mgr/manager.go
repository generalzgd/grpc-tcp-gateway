/**
 * @version: 1.0.0
 * @author: zhangguodong:general_zgd
 * @license: LGPL v3
 * @contact: general_zgd@163.com
 * @site: github.com/generalzgd
 * @software: GoLand
 * @file: manager.go
 * @time: 2019/9/23 15:08
 */
package mgr

import (
	`context`
	`crypto/tls`
	`encoding/json`
	`errors`
	`fmt`
	`io`
	`net`
	`net/http`
	`runtime/debug`
	`strings`
	`sync`
	`time`

	`github.com/astaxie/beego/logs`
	comm_libs `github.com/generalzgd/comm-libs`
	`github.com/generalzgd/grpc-svr-frame/common`
	`github.com/generalzgd/grpc-svr-frame/config/ymlcfg`
	_ `github.com/generalzgd/grpc-svr-frame/grpc-consul`
	ctrl `github.com/generalzgd/grpc-svr-frame/grpc-ctrl`
	`github.com/generalzgd/grpc-svr-frame/monitor`
	`github.com/generalzgd/grpc-svr-frame/monitor/analyse`
	_ `github.com/generalzgd/grpc-svr-frame/monitor/gorute`
	_ `github.com/generalzgd/grpc-svr-frame/monitor/mem`
	_ `github.com/generalzgd/grpc-svr-frame/monitor/tps`
	`github.com/generalzgd/grpc-svr-frame/prewarn`
	gwproto `github.com/generalzgd/grpc-tcp-gateway-proto/goproto`
	`github.com/generalzgd/link`
	`github.com/golang/protobuf/proto`
	`github.com/gorilla/websocket`
	grpcpool `github.com/processout/grpc-go-pool`
	`google.golang.org/grpc/metadata`

	`github.com/generalzgd/grpc-tcp-gateway/codec`
	`github.com/generalzgd/grpc-tcp-gateway/config`
)

var (
	mgrInst     *Manager
	mgrInstOnce sync.Once

	upgrader = websocket.Upgrader{
		ReadBufferSize:  4096,
		WriteBufferSize: 4096,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	//
	analyseFail = errors.New("parse data fail")
)

type Manager struct {
	ctrl.GrpcController
	cfg       config.AppConfig
	closeOnce sync.Once
	closeFlag bool
	wsCfg     config.TcpConnConfig
}

/*单例*/
func GetManagerInst() *Manager {
	if mgrInst == nil {
		mgrInstOnce.Do(func() {
			mgrInst = &Manager{}
			mgrInst.GrpcController = ctrl.MakeGrpcController()
		})
	}
	return mgrInst
}

func (p *Manager) Init(cfg config.AppConfig) {
	p.cfg = cfg

	prewarn.SetSendMailCallback(func(s string) {

	})
	// monitor.SetWarnHandler(prewarn.NewWarn)

	// 注册一个分析字段监控
	monitor.Register(analyse.NewAnalyse(1000, 10, monitor.ANALYSE_SUM, "", p.parseDataForMonitor))
}

func (p *Manager) Destroy() {
	p.closeOnce.Do(func() {
		p.closeFlag = true
	})
}

func (p *Manager) ServeClient() error {
	for _, it := range p.cfg.ServeList {
		switch it.Type {
		case common.GW_TYPE_TCP:
			if err := p.serveTcp(it); err != nil {
				return err
			}
		case common.GW_TYPE_WS:
			if err := p.serveWs(it); err != nil {
				return err
			}
		}
	}
	return nil
}

func (p *Manager) serveTcp(cfg config.TcpConnConfig) error {
	addr := fmt.Sprintf(":%d", cfg.Port)
	var server *link.Server
	if cfg.Secure {
		tlsCfg, err := p.GetTlsConfig(cfg.CertFiles...)
		if err != nil {
			return err
		}
		ln, err := net.Listen("tcp", addr)
		if err != nil {
			return err
		}
		lis := tls.NewListener(ln, tlsCfg)
		server = link.NewServer(
			lis,
			link.ProtocolFunc(func(rw io.ReadWriter) (link.Codec, error) {
				return codec.NewTlsCodec(rw.(net.Conn), cfg.BufferSize)
			}),
			cfg.SendChanSize,
			link.HandlerFunc(func(session *link.Session) {
				p.handleSession(session, cfg.MaxConn, cfg.IdleTimeout)
			}),
		)
	} else {
		lis, err := net.Listen("tcp", addr)
		if err != nil {
			return err
		}
		server = link.NewServer(
			lis,
			link.ProtocolFunc(func(rw io.ReadWriter) (link.Codec, error) {
				return codec.NewTcpCodec(rw.(net.Conn), cfg.BufferSize)
			}),
			cfg.SendChanSize,
			link.HandlerFunc(func(session *link.Session) {
				p.handleSession(session, cfg.MaxConn, cfg.IdleTimeout)
			}),
		)
	}

	go func() {
		if err := server.Serve(); err != nil {
			logs.Error("tcp gate way serve error.", err)
		}
	}()

	p.wsCfg = cfg
	logs.Info("start serve %s(%s) with secure(%v)", cfg.Name, addr, cfg.Secure)
	return nil
}

func (p *Manager) serveWs(cfg config.TcpConnConfig) error {
	var (
		httpServeMux = http.NewServeMux()
		err          error
	)
	httpServeMux.HandleFunc("/", p.serveWebSocket)
	httpServeMux.HandleFunc("/health", p.serveWebHealth)
	addr := fmt.Sprintf(":%d", cfg.Port)
	server := &http.Server{Addr: addr, Handler: httpServeMux}
	server.SetKeepAlivesEnabled(true)

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	if cfg.Secure {
		tlsCfg, err := p.GetTlsConfig(cfg.CertFiles...)
		if err != nil {
			return err
		}
		ln = tls.NewListener(ln, tlsCfg)
	}

	go func() {
		if err = server.Serve(ln); err != nil {
			logs.Info("start serve %s(%s) with secure(%s)", cfg.Name, addr, cfg.Secure)
		}
	}()

	logs.Info("start serve %s(%s) with secure(%v)", cfg.Name, addr, cfg.Secure)
	return nil
}

/*
* todo 响应阿里云slb的健康检测
 */
func (p *Manager) serveWebHealth(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(201)
}

/*
* todo websocket处理
 */
func (p *Manager) serveWebSocket(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(w, "Method Not Allowed", 405)
		return
	}
	ws, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		logs.Error("Websocket Upgrade error(%v), userAgent(%v)", err, req.UserAgent())
		return
	}

	var (
		lAddr = ws.LocalAddr()
		rAddr = ws.RemoteAddr()
	)
	// nginx代理的客户端ip获取
	realIp := req.Header.Get("X-Real-Ip")

	if len(realIp) == 0 {
		// slb的客户端ip获取。X-Forwarded-For: 用户真实IP, 代理服务器1-IP， 代理服务器2-IP，...
		tmp := req.Header.Get("X-Forwarded-For")
		splitIdx := strings.Index(tmp, ",")
		if splitIdx > 0 {
			realIp = tmp[:splitIdx]
		} else {
			realIp = tmp
		}
	}

	session := link.NewSession(codec.NewWssCodec(ws), p.wsCfg.SendChanSize)

	session.SetRealIp(realIp)

	logs.Debug("start websocket serve", lAddr, "with", rAddr, ">>", realIp)

	p.handleSession(session, p.wsCfg.MaxConn, p.wsCfg.IdleTimeout)
}

func (p *Manager) handleSession(session *link.Session, maxConn int, idleTimeout time.Duration) {
	conn := session.Codec()
	clientIp := session.GetRealIp()
	sid := conn.SocketID()

	info := &common.ClientConnInfo{
		SocketId: sid,
		ClientIp: clientIp,
		// GateIp:   p.gateIp,
	}
	session.State = info
	logs.Info("session connected ip:%s sid:%d", clientIp, sid)
	// 底层错误信息句柄
	// var err error
	defer func() {
		logs.Info("session disconnected ip:%s sid:%d", clientIp, sid)
		if r := recover(); r != nil {
			logs.Error("session panic ip:%s sid:%d, err:%v, stack:%s", clientIp, sid, r, string(debug.Stack()))
		}
		session.State = nil
	}()

	for {
		if idleTimeout > 0 {
			if err := conn.SetReadDeadline(time.Now().Add(idleTimeout)); err != nil {
				logs.Error("session timeout error. ip:%s sid:%d err:%v", clientIp, sid, err)
				return
			}
		}
		buf, err := session.Receive()
		if err != nil {
			logs.Error("session receive error:%v. ip:%s sid:%d err:%v", err, clientIp, sid, err)
			return
		}

		data := buf.([]byte)
		packet := codec.GateClientPack{}
		if err := codec.DecodePacket(data, &packet); err != nil {
			logs.Error("session decode packet error. ip:%s sid:%d err:%v", clientIp, sid, err)
			return
		}

		monitor.NewRecord(monitor.Stat_Tps)

		if packet.Id == codec.ID_Heartbeat {
			logs.Debug("session receive packet ip:%s sid:%d head:%v", clientIp, sid, packet.GateClientPackHead)
			continue
		}
		// todo 预处理其它逻辑

		// 	转发协议，转发失败则发送错误
		go func(session *link.Session, pack codec.GateClientPack, info *common.ClientConnInfo, sid uint32) {
			if err := p.transmitPack(session, packet, info); err != nil {
				logs.Error("session transmit pack error. ip:%s sid:%d err:%v", clientIp, sid, err)
			}
		}(session, packet, info, sid)
	}
}

// 转换协议并发送, 前提是解析出当前的包
func (p *Manager) transmitPack(session *link.Session, pack codec.GateClientPack, info *common.ClientConnInfo) (err error) {
	// begin := time.Now()
	defer func() {
		if r := recover(); r != nil {
			logs.Error("session process pack panic.", r, string(debug.Stack()))
		}
		// 异常返回
		if err != nil {
			imerr := &gwproto.GwError{
				Code:    1,
				Message: fmt.Sprintf("%v", err),
			}
			p.sendReplyPack(session, pack, imerr)
		}
	}()
	var md metadata.MD
	// 根据cmdid映射，得到对应的后端方法名称 package.Service/Method, 例如：ZQProto.Authorize/Login
	meth := gwproto.GetMethById(pack.Id)
	if len(meth) < 1 {
		err = codec.IdFieldError
		return
	}

	cfg, ok := p.getEndpointByMeth(meth)
	if !ok {
		err = codec.EndpointError
		return
	}

	// 用于统计分析
	monitor.NewRecord(monitor.Stat_Analyse, pack)

	// 转换用户链接信息为metadata.MD，用于grpc的header传输。
	// 接收方将header信息转换成ClientConnInfo结构体，以获得用户链接信息
	if md, err = p.ClientInfoToMD(info); err != nil {
		return
	}
	// grpc传输结束时的方法调用
	doneHandler := func(reply proto.Message) {
		p.sendReplyPack(session, pack, reply)
	}
	logs.Debug("start call grpc:", cfg.Address, fmt.Sprintf("head:%v", pack.GateClientPackHead))

	var conn *grpcpool.ClientConn
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	conn, err = p.GetGrpcConnWithLB(cfg, ctx)
	if err != nil {
		return
	}
	defer conn.Close()

	args := &gwproto.TransmitArgs{
		Method:       meth,
		Endpoint:     cfg.Address,
		Conn:         conn.ClientConn,
		MD:           md,
		Data:         pack.Body,
		Codec:        pack.Codec,
		DoneCallback: doneHandler,
		Opts:         nil,
	}
	// 将pack的信息，转换传输给后端的服务
	if err = gwproto.RegisterTransmitor(args); err != nil {
		return
	}
	return
}

// 发送响应包，包头沿用了旧包头，即复用了seq，opt字段
func (p *Manager) sendReplyPack(session *link.Session, pack codec.GateClientPack, reply proto.Message) {
	var bts []byte
	var err error
	if pack.Codec == codec.PACK_CODEC_PROTO {
		if bts, err = proto.Marshal(reply); err != nil {
			return
		}
	} else if pack.Codec == codec.PACK_CODEC_JSON {
		if bts, err = json.Marshal(reply); err == nil {
			return
		}
	}

	pack.Id = gwproto.GetIdByMsgObj(reply)
	pack.Length = uint16(len(bts))
	pack.Body = bts
	p.SendToClient(session, codec.EncodePacket(&pack))
}

// 统一封装session 发送方法
func (p *Manager) SendToClient(session *link.Session, msg []byte) error {
	err := session.Send(msg)
	if err != nil {
		session.Close()
	}
	// 下行也要记录tps
	monitor.NewRecord(monitor.Stat_Tps)
	return err
}

func (p *Manager) getEndpointByMeth(meth string) (ymlcfg.EndpointConfig, bool) {
	_, tarSvr, _, _ := gwproto.ParseMethod(meth)
	tarSvr = strings.ToLower(tarSvr)
	cfg, ok := p.cfg.EndpointSvr[tarSvr]
	return cfg, ok
}

// 解析对应的字段
func (p *Manager) parseDataForMonitor(data interface{}, field string) (int, error) {
	if data == nil || field == "" {
		return 0, analyseFail
	}
	pack, ok := data.(codec.GateClientPack)
	if !ok {
		return 0, analyseFail
	}
	obj := gwproto.GetMsgObjById(pack.Id)
	if obj == nil {
		return 0, analyseFail
	}

	if err := proto.Unmarshal(pack.Body, obj); err != nil {
		return 0, analyseFail
	}

	res := comm_libs.GetFieldValueFromTarget(field, obj)
	if res == nil {
		return 0, analyseFail
	}
	return comm_libs.Interface2Int(res), nil
}
