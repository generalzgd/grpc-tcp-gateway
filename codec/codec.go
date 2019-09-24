/**
 * @version: 1.0.0
 * @author: zhangguodong:general_zgd
 * @license: LGPL v3
 * @contact: general_zgd@163.com
 * @site: github.com/generalzgd
 * @software: GoLand
 * @file: codec.go
 * @time: 2019/9/23 15:08
 */
package codec

import (
	`bufio`
	`encoding/binary`
	`io`
	`net`
	`sync/atomic`
	`time`

	`github.com/generalzgd/link`
	`github.com/gorilla/websocket`
)

var (
	connSeed = uint32(0)
)

func NewTlsCodec(conn net.Conn, bufferSize int) (link.Codec, error) {
	id := atomic.AddUint32(&connSeed, 1)
	c := &TlsGateCodec{
		GateCodecBase: GateCodecBase{
			id:                    id,
			conn:                  conn,
			reader:                bufio.NewReaderSize(conn, bufferSize),
		},
	}
	c.headBuf = c.headDat[:]
	return c, nil
}

func NewTcpCodec(conn net.Conn, bufferSize int) (link.Codec, error) {
	id := atomic.AddUint32(&connSeed, 1)
	c := &TcpGateCodec{
		GateCodecBase: GateCodecBase{
			id:                    id,
			conn:                  conn,
			reader:                bufio.NewReaderSize(conn, bufferSize),
		},
	}
	c.headBuf = c.headDat[:]
	return c, nil
}

func NewWssCodec(conn *websocket.Conn) *WsGateCodec {
	id := atomic.AddUint32(&connSeed, 1)
	c := &WsGateCodec{
		conn:                  conn,
	}
	c.id = id
	c.headBuf = c.headDat[:]
	return c
}



// ////////
type GateCodecBase struct {
	id      uint32
	conn    net.Conn
	reader  *bufio.Reader
	headBuf []byte
	headDat [PACK_HEAD_SIZE]byte
}

func (p *GateCodecBase) MemSet(tar []byte, val byte) {
	for i := range tar {
		tar[i] = val
	}
}

type TlsGateCodec struct {
	GateCodecBase
	realIp string
}

func (p *TlsGateCodec) ClearSendChan(sendChan <-chan interface{}) {
}

func (p *TlsGateCodec) SocketID() uint32 {
	return p.id
}

func (p *TlsGateCodec) ClientAddr() string {
	return p.conn.RemoteAddr().String()
}

func (p *TlsGateCodec) SetReadDeadline(t time.Time) error {
	return p.conn.SetReadDeadline(t)
}

func (p *TlsGateCodec) Receive() (interface{}, error) {
	p.MemSet(p.headBuf, 0)
	if _, err := io.ReadFull(p.reader, p.headBuf); err != nil {
		return nil, err
	}
	length := int(binary.LittleEndian.Uint16(p.headBuf[0:]))
	if length > PACK_MAX_SIZE {
		return nil, TooLargeError
	}
	buff := make([]byte, PACK_HEAD_SIZE + length)
	copy(buff, p.headBuf)
	if _, err := io.ReadFull(p.reader, buff[PACK_HEAD_SIZE:]); err != nil {
		return nil, err
	}
	return buff, nil
}

func (p *TlsGateCodec) Send(msg interface{}) error {
	buffer := msg.([]byte)
	_, err := p.conn.Write(buffer)
	return err
}

func (p *TlsGateCodec) Close() error {
	if p.conn != nil {
		return p.conn.Close()
	}
	return nil
}

// ////
type TcpGateCodec struct {
	GateCodecBase
	realIp string
}

func (p *TcpGateCodec) ClearSendChan(sendChan <-chan interface{}) {
}

func (p *TcpGateCodec) SocketID() uint32 {
	return p.id
}

func (p *TcpGateCodec) ClientAddr() string {
	return p.conn.RemoteAddr().String()
}

func (p *TcpGateCodec) SetReadDeadline(t time.Time) error {
	return p.conn.SetReadDeadline(t)
}

func (p *TcpGateCodec) Receive() (interface{}, error) {
	p.MemSet(p.headBuf, 0)
	if _, err := io.ReadFull(p.reader, p.headBuf); err != nil {
		return nil, err
	}
	length := int(binary.LittleEndian.Uint16(p.headBuf[0:]))
	if length > PACK_MAX_SIZE {
		return nil, TooLargeError
	}
	buff := make([]byte, PACK_HEAD_SIZE + length)
	copy(buff, p.headBuf)

	if _, err := io.ReadFull(p.reader, buff[PACK_HEAD_SIZE:]); err != nil {
		return nil, err
	}
	// logs.Debug("head %v, receive %v",  p.headBuf, buff[p.HeadSize:])
	return buff, nil
}

func (p *TcpGateCodec) Send(msg interface{}) error {
	buffer := msg.([]byte)
	_, err := p.conn.Write(buffer)
	return err
}

func (p *TcpGateCodec) Close() error {
	if p.conn != nil {
		return p.conn.Close()
	}
	return nil
}

// *********************************************
type WsGateCodec struct {
	GateCodecBase
	// id      uint32
	conn    *websocket.Conn
	// headBuf []byte
	// headDat [PACK_HEAD_SIZE]byte
	realIp  string
}

func (p *WsGateCodec) Receive() (interface{}, error) {
	p.MemSet(p.headBuf, 0)

	_, r, err := p.conn.NextReader()
	if err != nil {
		return nil, err
	}
	if _, err := io.ReadFull(r, p.headBuf); err != nil {
		return nil, err
	}
	length := int(binary.LittleEndian.Uint16(p.headBuf[0:4]))
	if length > PACK_MAX_SIZE {
		return nil, TooLargeError
	}

	buffer := make([]byte, PACK_HEAD_SIZE + length)
	copy(buffer, p.headBuf)
	if _, err := io.ReadFull(r, buffer[PACK_HEAD_SIZE:]); err != nil {
		return nil, err
	}
	return buffer, nil
}

func (p *WsGateCodec) Send(msg interface{}) error {
	buffer := msg.([]byte)
	err := p.conn.WriteMessage(websocket.BinaryMessage, buffer) // p.conn.Write(buffer)
	return err
}

func (p *WsGateCodec) Close() error {
	if p.conn != nil {
		return p.conn.Close()
	}
	return nil
}

func (p *WsGateCodec) SocketID() uint32 {
	return p.id
}

func (p *WsGateCodec) ClientAddr() string {
	return p.conn.RemoteAddr().String()
}

func (p *WsGateCodec) SetReadDeadline(t time.Time) error {
	return p.conn.SetReadDeadline(t)
}
