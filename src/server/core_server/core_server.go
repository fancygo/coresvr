package main

//服务主程序

import (
	"frame"
	"frame/conn"
	"frame/def"
	"frame/logger"
	"frame/svr"
	"net"
	"server/core_server/logic/msgproc"
	"sync"
)

func main() {
	//初始化主服务
	svrId := def.CORE_SVR_ID
	coreSvr := svr.NewSvr(svrId)
	coreSvr.Init()

	addr := frame.GetSvrIP(svrId) + ":" + frame.GetSvrPort(svrId)
	logger.Info("%s svr addr = %+v", frame.GetSvrName(svrId), addr)
	tcp_addr, err := net.ResolveTCPAddr("tcp", addr)
	frame.CheckErr(err)
	listener, err := net.ListenTCP("tcp", tcp_addr)
	frame.CheckErr(err)

	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(1)
	go doListen(listener, waitGroup)
	waitGroup.Wait()
}

func doListen(listener *net.TCPListener, waitGroup *sync.WaitGroup) {
	defer func() {
		listener.Close()
		waitGroup.Done()
	}()
	//监听svr_cli连接
	for {
		logger.Debug("accept wait")
		connTcp, err := listener.AcceptTCP()
		if err != nil {
			logger.Debug("listener accept err = %+v", err)
			continue
		}
		logger.Debug("connTcp = %+v, %s", connTcp.RemoteAddr().Network(), connTcp.RemoteAddr().String())
		cliConn := conn.NewCliConn(connTcp)
		cliConn.RegFunc(msgproc.ProcMsgData)
		//conn.Add(cliConn)
		go cliConn.HandleClient(nil)
	}
}
