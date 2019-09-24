/**
 * @version: 1.0.0
 * @author: zhangguodong:general_zgd
 * @license: LGPL v3
 * @contact: general_zgd@163.com
 * @site: github.com/generalzgd
 * @software: GoLand
 * @file: main.go
 * @time: 2019/9/23 14:58
 */
package main

import (
	`os`
	`os/signal`
	`runtime`
	`syscall`

	`github.com/astaxie/beego/logs`

	`github.com/generalzgd/grpc-tcp-gateway/config`
	`github.com/generalzgd/grpc-tcp-gateway/mgr`
)

func init() {
	logger := logs.GetBeeLogger()
	logger.SetLevel(logs.LevelInfo)
	logger.SetLogger(logs.AdapterConsole)
	logger.SetLogger(logs.AdapterFile, `{"filename":"logs/file.log","level":7,"maxlines":1024000000,"maxsize":1024000000,"daily":true,"maxdays":7}`)
	logger.EnableFuncCallDepth(true)
	logger.SetLogFuncCallDepth(4)
	logger.Async(100000)
}

func exit(err error) {
	code := 0
	if err != nil {
		logs.Error("got error:%v", err)
		logs.GetBeeLogger().Flush()
	}
	os.Exit(code)
}

func main() {
	var err error
	defer func() {
		exit(err)
	}()

	runtime.GOMAXPROCS(runtime.NumCPU())
	// do something
	cfg := config.AppConfig{}
	if err = cfg.Load(); err != nil {
		return
	}

	manager := mgr.GetManagerInst()
	manager.Init(cfg)

	if err = manager.ServeClient(); err != nil {
		return
	}

	// catchs system signal
	chSig := make(chan os.Signal)
	signal.Notify(chSig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGTERM)
	sig := <-chSig
	logs.Info("siginal:", sig)

	manager.Destroy()
}
