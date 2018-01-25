package conn

import (
	"sync"
)

type ConnItf interface {
	GetId() uint32
	GerAddr() string
	SendChanData(data []byte)
	DoSend(chStop <-chan bool)
	RecvChanData(data []byte)
	DoRecv(chStop <-chan bool)
	procMsgData(data []byte)
	HandleClient(wait_group *sync.WaitGroup)
}
