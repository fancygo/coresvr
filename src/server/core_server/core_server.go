package main

//服务主程序

import (
	"frame"
	"frame/logger"
	"frame/svr"
	"net"
)

func main() {
	//初始化主服务
	svrId := frame.GetMainSvr()
	coreSvr := svr.NewSvr(svrId)
	coreSvr.Init()

	addr := frame.GetSvrIP(svrId) + ":" + frame.GetSvrPort(svrId)
	logger.Info("%s svr addr = %+v", frame.GetSvrName(svrId), addr)
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
	}
	//cliConn := conn.NewCliConn(connTcp)
	//conn.Add(cliConn)
	//go cliConn.HandleClient()
}
