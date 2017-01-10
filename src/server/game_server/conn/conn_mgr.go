package conn

//每个客户端连接实例管理器, 实现了连接注册, 卸载
//当时主要是为了实现广播和消息通知的作用

import (
	"frame/logger"
	"sync"
)

var (
	Conn map[uint32]*CliConn
	lock *sync.RWMutex
)

func init() {
	Conn = make(map[uint32]*CliConn)
	lock = new(sync.RWMutex)
}

func Add(cliConn *CliConn) {
	cid := cliConn.GetId()
	Conn[cid] = cliConn
	logger.Info("add conn cid = %d", cid)
}

func Del(cid uint32) {
	delete(Conn, cid)
	logger.Info("del conn cid = %d", cid)
}

//广播
func Broadcast(send_data []byte) {
	for _, v := range Conn {
		v.SendChanData(send_data)
	}
}
