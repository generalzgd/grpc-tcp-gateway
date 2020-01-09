/**
 * @version: 1.0.0
 * @author: zhangguodong:general_zgd
 * @license: LGPL v3
 * @contact: general_zgd@163.com
 * @site: github.com/generalzgd
 * @software: GoLand
 * @file: config.go
 * @time: 2019/9/23 15:09
 */
package config

import (
	`fmt`
	`os`
	`path/filepath`
	`time`

	`github.com/astaxie/beego/logs`
	`github.com/generalzgd/grpc-svr-frame/config/ymlcfg`
)

// 网关连接配置
type TcpConnConfig struct {
	Name            string            `yaml:"name"`
	Type            string            `yaml:"type"`      // http/tcp/ws/grpc
	Secure          bool              `yaml:"secure"`    // false: http/tcp/ws  true: https/tls/wss, 版本号默认1.1
	CertFiles       []ymlcfg.CertFile `yaml:"certfiles"` // 证书文件，pem格式
	BufferSize      int               `yaml:"buffersize"`
	MaxConn         int               `yaml:"maxconn"`
	IdleTimeout     time.Duration     `yaml:"idletimeout"`
	SendChanSize    int               `yaml:"sendchansize"` //
	ReceiveChanSize int               `yaml:"recvchansize"` //
	Port            uint32            `yaml:"port"`         // 侦听端口
}

type AppConfig struct {
	Name        string                           `yaml:"name"`
	Ver         string                           `yaml:"ver"`
	LogLevel    int                              `yaml:"loglevel"`
	Pprof       int                              `yaml:"pprof"`
	Consul      ymlcfg.ConsulConfig              `yaml:"consul"`
	EndpointSvr map[string]ymlcfg.EndpointConfig `yaml:"endpoint"` // key 对应tcpgate.proto文件中的@target标签
	ServeList   []TcpConnConfig                  `yaml:"servelist"`
	RpcSvr      ymlcfg.EndpointConfig            `yaml:"rpc"` // 下行，grpc 端口信息
}

func (p *AppConfig) GetLogLevel() int {
	if p.LogLevel <= 0 {
		return logs.LevelInfo
	}
	return p.LogLevel
}

func (p *AppConfig) Load() error {
	path := fmt.Sprintf("%s/config/cfg.yml", filepath.Dir(os.Args[0]))
	if err := ymlcfg.LoadYaml(path, p); err != nil {
		return err
	}
	return nil
}
