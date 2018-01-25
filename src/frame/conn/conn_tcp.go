package conn

import (
	_ "frame/def"
	"frame/logger"
	"net"
	"sync"
	"sync/atomic"
	_ "time"
)

const (
	maxUint32 = ^uint32(0)
	chanLen   = 15
)

var (
	gConnId = uint32(0) // 全局连接ID
)

type ProcFunc func([]byte)

type CliConn struct {
	cid   uint32
	addr  string
	accid string

	chanRecv chan []byte
	chanSend chan []byte

	conn *net.TCPConn

	procFunc ProcFunc
}

func NewCliConn(conn *net.TCPConn) *CliConn {
	atomic.AddUint32(&gConnId, 1)
	if gConnId >= maxUint32 {
		atomic.StoreUint32(&gConnId, 1)
	}

	return &CliConn{
		cid:      gConnId,
		conn:     conn,
		addr:     conn.RemoteAddr().String(),
		chanRecv: make(chan []byte, chanLen),
		chanSend: make(chan []byte, chanLen),
	}
}

func (this *CliConn) RegFunc(fun ProcFunc) {
	this.procFunc = fun
}

func (this *CliConn) GetId() uint32 {
	return this.cid
}

func (this *CliConn) GetAddr() string {
	return this.addr
}

func (this *CliConn) SendChanData(data []byte) {
	this.chanSend <- data
}

func (this *CliConn) DoSend(chStop <-chan bool) {
	for {
		select {
		case <-chStop:
			return
		case data := <-this.chanSend:
			this.conn.Write(data)
		}
	}
}

func (this *CliConn) RecvChanData(data []byte) {
	this.chanRecv <- data
}

func (this *CliConn) DoRecv(chStop <-chan bool) {
	for {
		select {
		case <-chStop:
			return
		case data := <-this.chanRecv:
			this.procFunc(data)
		}
	}
}

func (this *CliConn) HandleClient(wait_group *sync.WaitGroup) {
	chStop := make(chan bool)
	defer func() {
		logger.Info("conn id = %d, close", this.cid)
		//module.UserApi.DelUser(this.accid)
		//Del(this.cid)
		this.conn.Close()
		close(chStop)
	}()

	go this.DoSend(chStop)
	go this.DoRecv(chStop)

	for {
		buf := make([]byte, 256, 256)
		readLen, err := this.conn.Read(buf)
		if err != nil {
			return
		}

		logger.Info("read somsthing len = %d", readLen)
		if readLen <= 17 {
			logger.Error("recv msg err ,len < 17")
			continue
		}
		realData := make([]byte, readLen, readLen)
		copy(realData, buf[0:readLen])
		this.RecvChanData(realData)
	}
}
