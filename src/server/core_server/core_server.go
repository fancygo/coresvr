package main

//服务主程序

import (
	"flag"
	"frame"
	"frame/logger"
	"net"
	_ "time"
)

func main() {
	logger.Debug("-----------------------core server start-----------------------")
	//命令行
	flagConsole := flag.Bool("c", false, "show log in console")
	flag.Parse()
	if *flagConsole == true {
		logger.SetConsole(*flagConsole)
	}
	//初始化主服务
	svrId := frame.GetMainSvr()
	addr := frame.GetSvrIP(svrId) + ":" + frame.GetSvrPort(svrId)
	logger.Info("core svr addr = %+v", addr)
	tcp_addr, err := net.ResolveTCPAddr("tcp", addr)
	frame.CheckErr(err)
	listener, err := net.ListenTCP("tcp", tcp_addr)
	frame.CheckErr(err)

	//监听cli连接
	for {
		logger.Debug("accept wait")
		connTcp, err := listener.AcceptTCP()
		if err != nil {
			logger.Debug("listener accept err = %+v", err)
			continue
		}

		logger.Debug("connTcp = %+v, %s", connTcp.RemoteAddr().Network(), connTcp.RemoteAddr().String())
		//cliConn := conn.NewCliConn(connTcp)
		//conn.Add(cliConn)
		//go cliConn.HandleClient()
	}
}
