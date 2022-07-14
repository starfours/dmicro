package main

import (
	"github.com/gogf/gf/os/gtime"
	"github.com/osgochina/dmicro/dserver"
	"github.com/osgochina/dmicro/examples/dserver/sandbox"
	"github.com/osgochina/dmicro/logger"
	"time"
)

func main() {
	//dserver.CloseCtrl()
	dserver.SetName("DMicro")
	dserver.Setup(func(svr *dserver.DServer) {
		svr.ProcessModel(dserver.ProcessModelMulti)
		svr.SetInheritListener([]dserver.InheritAddr{
			{Network: "tcp", Host: "127.0.0.1", Port: 8199, ServerName: "default"},
			{Network: "http", Host: "127.0.0.1", Port: 8080, ServerName: "ghttp1"},
		})
		err := svr.AddSandBox(new(sandbox.DefaultSandBox), svr.NewService("admin"))
		if err != nil {
			logger.Fatal(err)
		}
		err = svr.AddSandBox(new(sandbox.DefaultSandBox1), svr.NewService("user"))
		if err != nil {
			logger.Fatal(err)
		}
		go func() {

			t := time.NewTicker(1 * time.Second)
			for {
				select {
				case <-t.C:
					logger.Infof("定时器打印日志：%s", gtime.Now().String())
				}
			}
		}()
	})
}
