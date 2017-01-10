package conn

//每个客户端实例
//包括id, 消息发送接收队列, 解析, 处理

import (
	"fmt"
	"frame/def"
	"frame/logger"
	"golang.org/x/net/websocket"
	module "server/game_server/game"
	"server/game_server/msgprocess"
	"sync"
	"sync/atomic"
	"time"
)

const (
	maxUint32 = ^uint32(0)
	chanLen   = 15
)

var (
	gConnId = uint32(0) // 全局连接ID
)

type CliConn struct {
	cid   uint32
	conn  *websocket.Conn
	addr  string
	accid string

	chanRecv chan []byte
	chanSend chan []byte
}

func NewCliConn(conn *websocket.Conn) *CliConn {
	//连接id的递增
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
			this.procMsgData(data)
		}
	}
}

func (this *CliConn) procMsgData(data []byte) {
	logger.Info("--------------Msg proc start------------")
	fmt.Println(data)
	start := time.Now()

	dataLen := len(data)
	//获取命令id
	msgidArray := make([]byte, 4, 4)
	copy(msgidArray, data[13:17])
	msgId := msgprocess.GetMsgID(msgidArray)
	logger.Info("recv msgid = %d", msgId)

	//真实消息数据
	msgData := data[17:dataLen]

	var sendData []byte
	var notiData []byte

	//根据消息号送到不同的消息处理函数
	switch msgId {
	case uint(def.MsgId_UserLogin):
		sendData, this.accid = msgprocess.UserLoginMsg(msgData)
	case uint(def.MsgId_UploadCraft):
		sendData, notiData = msgprocess.UploadCraftMsg(msgData)
	case uint(def.MsgId_GetCraft):
		sendData = msgprocess.GetCraftMsg(msgData)
	case uint(def.MsgId_Praise):
		sendData = msgprocess.PraiseCraftMsg(msgData)
	case uint(def.MsgId_GetPraise):
		sendData = msgprocess.GetPraiseMsg(msgData)
	case uint(def.MsgId_SyncInfo):
		sendData = msgprocess.SyncInfoMsg(msgData, this.accid)
	case uint(def.MsgId_GetPopular):
		sendData = msgprocess.GetPopularMsg(msgData)
	default:
		logger.Error("error data = %+v", data)
		return
	}
	if len(sendData) > 0 {
		this.SendChanData(sendData)
	}
	if len(notiData) > 0 {
		Broadcast(notiData)
	}

	secs := time.Since(start).Seconds()
	//记录每条消息的处理时间
	logger.Info("--------------Msg end last secs = %.5f-----------------", secs)
}

//每个连接实例的处理函数
func (this *CliConn) HandleClient(wait_group *sync.WaitGroup) {
	chStop := make(chan bool)
	defer func() {
		module.UserApi.DelUser(this.accid)
		Del(this.cid)
		this.conn.Close()
		close(chStop)
		wait_group.Done()
	}()

	go this.DoSend(chStop)
	go this.DoRecv(chStop)

	for {
		buf := make([]byte, 1024, 1024)
		readLen, err := this.conn.Read(buf)
		if err != nil {
			return
		}

		logger.Info("read somsthing len = %d", readLen)
		if readLen <= 17 {
			logger.Error("recv msg err, len < 17")
			continue
		}
		realData := make([]byte, readLen, readLen)
		copy(realData, buf[0:readLen])
		this.RecvChanData(realData)
	}
}
