package msgprocess

//这边进行消息的解包和打包
//实际的处理放到msg当中

import (
	"encoding/json"
	"frame"
	"frame/def"
	"frame/logger"
)

func constructMsg(commandid uint, msglen int) []byte {
	sendData := make([]byte, 9, 9)
	lengthArray := make([]byte, 4, 4)
	var length uint = uint(msglen + 4)
	lengthArray = frame.IntToArray(length)

	sendCommandArray := make([]byte, 4, 4)
	var command uint = commandid
	sendCommandArray = frame.IntToArray(command)

	sendData = append(sendData, lengthArray[0:]...)
	sendData = append(sendData, sendCommandArray[0:]...)
	//logger.Info("sendData len = %d\n", len(sendData))
	return sendData
}

func GetMsgID(msgid_array []byte) uint {
	return frame.ArrayToInt(msgid_array)
}

func GetMsgLen(len_array []byte) uint {
	return frame.ArrayToInt(len_array)
}

func UserLoginMsg(recv_data []byte) ([]byte, string) {
	sendData := make([]byte, 0)
	recvMsg := &def.UserLoginA{}
	sendMsg := &def.UserLoginR{
		Ret:  def.OK,
		Info: &def.UserInfo{},
	}
	logger.Info("recv_data = %s", string(recv_data))
	err := json.Unmarshal(recv_data, recvMsg)
	if err != nil {
		logger.Error("UserVerify json unmarshal err = %+v", err)
		return sendData, ""
	}
	accid := UserLogin(recvMsg, sendMsg)

	sendMsgData, err := json.Marshal(sendMsg)
	if err != nil {
		logger.Error("UserVerify json marshal err = %+v", err)
		return sendData, ""
	}
	msgLen := len(sendMsgData)
	sendData = constructMsg(uint(def.MsgId_UserLoginRet), msgLen)
	sendData = append(sendData, sendMsgData...)
	//sendData = append(sendData, byte('0'))

	DumpMsgInfo(int(def.MsgId_UserLoginRet), sendMsgData, msgLen, sendData, len(sendData))
	return sendData, accid
}

func UploadCraftMsg(recv_data []byte) ([]byte, []byte) {
	sendData := make([]byte, 0)
	notiData := make([]byte, 0)
	sendMsg := &def.UploadCraftR{}
	author := new(string)
	ifbroad := new(int)
	UploadCraft(recv_data, sendMsg, author, ifbroad)

	sendMsgData, err := json.Marshal(sendMsg)
	if err != nil {
		logger.Error("UploadNewCraftRet json marshal err = %+v", err)
		return sendData, notiData
	}
	msgLen := len(sendMsgData)
	sendData = constructMsg(uint(def.MsgId_UploadCraftRet), msgLen)
	sendData = append(sendData, sendMsgData...)
	DumpMsgInfo(int(def.MsgId_UploadCraftRet), sendMsgData, msgLen, sendData, len(sendData))

	//通知
	notiMsgData := make([]byte, 0)
	if *ifbroad == 1 {
		notiMsg := &def.UploadCraftNoti{
			Author: *author,
		}
		notiMsgData, _ = json.Marshal(notiMsg)
		msgLen := len(notiMsgData)
		notiData = constructMsg(uint(def.MsgId_UploadCraftNoti), msgLen)
		notiData = append(notiData, notiMsgData...)
		DumpMsgInfo(int(def.MsgId_UploadCraftNoti), notiMsgData, msgLen, notiData, len(notiData))
	}
	return sendData, notiData
}

func GetCraftMsg(recv_data []byte) []byte {
	sendData := make([]byte, 0)
	sendMsg := &def.GetCraftR{}
	GetCraft(recv_data, sendMsg)

	sendMsgData, err := json.Marshal(sendMsg)
	if err != nil {
		logger.Error("GetCraftRet json marshal err = %+v", err)
		return sendData
	}
	msgLen := len(sendMsgData)
	sendData = constructMsg(uint(def.MsgId_GetCraftRet), msgLen)
	sendData = append(sendData, sendMsgData...)
	//sendData = append(sendData, "")

	DumpMsgInfo(int(def.MsgId_GetCraftRet), sendMsgData, msgLen, sendData, len(sendData))
	return sendData
}

func PraiseCraftMsg(recv_data []byte) []byte {
	sendData := make([]byte, 0)
	sendMsg := &def.PraiseCraftR{}
	PraiseCraft(recv_data, sendMsg)

	sendMsgData, err := json.Marshal(sendMsg)
	if err != nil {
		logger.Error("PraiseCraftRet json marshal err = %+v", err)
		return sendData
	}
	msgLen := len(sendMsgData)
	sendData = constructMsg(uint(def.MsgId_PraiseRet), msgLen)
	sendData = append(sendData, sendMsgData...)
	//sendData = append(sendData, "")

	DumpMsgInfo(def.MsgId_PraiseRet, sendMsgData, msgLen, sendData, len(sendData))
	return sendData
}

func GetPraiseMsg(recv_data []byte) []byte {
	sendData := make([]byte, 0)
	sendMsg := &def.GetPraiseR{}
	GetPraise(recv_data, sendMsg)

	sendMsgData, err := json.Marshal(sendMsg)
	if err != nil {
		logger.Error("GetPraiseMsgRet json marshal err = %+v", err)
		return sendData
	}
	msgLen := len(sendMsgData)
	sendData = constructMsg(uint(def.MsgId_GetPraiseRet), msgLen)
	sendData = append(sendData, sendMsgData...)
	//sendData = append(sendData, "")

	DumpMsgInfo(def.MsgId_GetPraiseRet, sendMsgData, msgLen, sendData, len(sendData))
	return sendData
}

func GetPopularMsg(recv_data []byte) []byte {
	sendData := make([]byte, 0)
	sendMsg := &def.GetPopularR{}
	GetPopular(recv_data, sendMsg)

	sendMsgData, err := json.Marshal(sendMsg)
	if err != nil {
		logger.Error("GetPopularRet json marshal err = %+v", err)
		return sendData
	}
	msgLen := len(sendMsgData)
	sendData = constructMsg(uint(def.MsgId_GetPopularRet), msgLen)
	sendData = append(sendData, sendMsgData...)
	DumpMsgInfo(def.MsgId_GetPopularRet, sendMsgData, msgLen, sendData, len(sendData))
	return sendData
}

func SyncInfoMsg(recv_data []byte, accid string) []byte {
	sendData := make([]byte, 0)
	sendMsg := &def.SyncInfoR{}
	SyncInfo(recv_data, sendMsg, accid)
	return sendData
}

func DumpMsgInfo(msgid int, msg interface{}, msg_len interface{}, send_data interface{}, send_len interface{}) {
	logger.Info("send msgid = %d", msgid)
	logger.Info("send MsgData = %s", string(msg.([]byte)))
	logger.Info("msg Len = %d", msg_len)
	//logger.Info("send Data = %v", send_data)
	logger.Info("send Len = %d", send_len)
}
