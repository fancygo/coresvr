package msgprocess

//把解包后的实际数据送到游戏模块中处理, 把处理完的数据序列化返回

import (
	"encoding/json"
	"frame/def"
	"frame/logger"
	_ "github.com/ugorji/go/codec"
	module "server/game_server/game"
)

func UserLogin(recv_msg *def.UserLoginA, send_msg *def.UserLoginR) string {
	accid := recv_msg.Accid
	name := recv_msg.Name

	if len(accid) == 0 || len(name) == 0 {
		send_msg.Ret = def.Ret_LoginAccidNameNUll
		logger.Info("User Login accid = %s, name = %s", accid, name)
		return ""
	}
	//创建新账号
	user := module.UserApi.InitNewUser(accid, name)
	send_msg.Ret = def.OK
	json.Unmarshal(([]byte)(user.Info), send_msg.Info)
	//加入管理器
	module.UserApi.AddUser(user)
	return user.Accid
}

func UploadCraft(recv_data []byte, send_msg *def.UploadCraftR, author *string, ifbroad *int) {
	send_msg.Ret = def.OK

	recvMsg := &def.UploadCraftA{}
	logger.Info("recv_data = %s", string(recv_data))
	err := json.Unmarshal(recv_data, recvMsg)
	if err != nil {
		logger.Error("UploadCraft json unmarshal err = %+v", err)
		send_msg.Ret = def.Ret_UnMarshalERR
		return
	}

	craftTyp := recvMsg.Typ
	craft := recvMsg.Craft
	craftAuthor := recvMsg.Author
	craftRect := craft.Rect
	craftData := craft.Data
	*author = recvMsg.Author
	*ifbroad = recvMsg.Ifbroad

	if craftRect == 0 {
		send_msg.Ret = def.Ret_CraftRectErr
	} else {
		if craftTyp == def.CraftGood {
			id := module.CraftApi.AddGoodCraft(craftAuthor, craftRect, craftData)
			send_msg.Craftid = id
		} else if craftTyp == def.CraftNormal {
			id := module.CraftApi.AddNormalCraft(craftAuthor, craftRect, craftData)
			send_msg.Craftid = id
		} else {
			send_msg.Ret = def.Ret_CraftTypeErr
		}
	}
	return
}

func GetCraft(recv_data []byte, send_msg *def.GetCraftR) {
	send_msg.Ret = def.OK

	recvMsg := &def.GetCraftA{}
	logger.Info("recv_data = %s", string(recv_data))
	err := json.Unmarshal(recv_data, recvMsg)
	if err != nil {
		logger.Error("GetCraft json unmarshal err = %+v", err)
		send_msg.Ret = def.Ret_UnMarshalERR
		return
	}

	var allCraft []*module.Craft
	idx := recvMsg.Idx

	craftTyp := recvMsg.Typ
	if craftTyp == def.CraftGood {
		allCraft = module.CraftApi.GetGoodCraft()
	} else if craftTyp == def.CraftNormal {
		allCraft = module.CraftApi.GetNormalCraft()
	} else {
		send_msg.Ret = def.Ret_CraftRectErr
		return
	}

	num := 0
	allLen := len(allCraft)
	if idx < allLen {
		if idx+def.GetCraftNum < allLen {
			num = def.GetCraftNum
		} else {
			num = allLen - idx
		}
		send_msg.Crafts = make([]*def.CraftEntry, 0, num)
		for i := idx; i < idx+num; i++ {
			craft := allCraft[i]
			craftEntry := &def.CraftEntry{
				Id:     craft.Id,
				Rect:   craft.Rect,
				Author: craft.Author,
				Data:   craft.Data,
				Praise: craft.Praise,
			}
			send_msg.Crafts = append(send_msg.Crafts, craftEntry)
		}
	}

	return
}

func PraiseCraft(recv_data []byte, send_msg *def.PraiseCraftR) {
	send_msg.Ret = def.OK

	recvMsg := &def.PraiseCraftA{}
	logger.Info("recv_data = %s", string(recv_data))
	err := json.Unmarshal(recv_data, recvMsg)
	if err != nil {
		logger.Error("PraiseCraft json unmarshal err = %+v", err)
		send_msg.Ret = def.Ret_UnMarshalERR
		return
	}

	id := recvMsg.Id

	send_msg.Ret = module.CraftApi.PraisePopular(id)
	return
}

func GetPraise(recv_data []byte, send_msg *def.GetPraiseR) {
	send_msg.Ret = def.OK

	recvMsg := &def.GetPraiseA{}
	logger.Info("recv_data = %s", string(recv_data))
	err := json.Unmarshal(recv_data, recvMsg)
	if err != nil {
		logger.Error("GetCraft json unmarshal err = %+v", err)
		send_msg.Ret = def.Ret_UnMarshalERR
		return
	}

	var allCraft []*module.Craft
	idx := recvMsg.Idx
	logger.Info("GetCraft idx = %d", idx)
	craftTyp := recvMsg.Typ

	if craftTyp == def.CraftGood {
		allCraft = module.CraftApi.GetGoodCraft()
	} else if craftTyp == def.CraftNormal {
		allCraft = module.CraftApi.GetNormalCraft()
	} else {
		send_msg.Ret = def.Ret_CraftRectErr
		return
	}

	num := 0
	allLen := len(allCraft)
	if idx < allLen {
		if idx+def.GetPraiseNum < allLen {
			num = def.GetPraiseNum
		} else {
			num = allLen - idx
		}
		send_msg.Praises = make([]*def.PraiseEntry, 0, num)
		for i := idx; i < idx+num; i++ {
			craft := allCraft[i]
			praiseEntry := &def.PraiseEntry{
				Id:     craft.Id,
				Praise: craft.Praise,
			}
			send_msg.Praises = append(send_msg.Praises, praiseEntry)
		}
	}
	return
}

func GetPopular(recv_data []byte, send_msg *def.GetPopularR) {
	send_msg.Ret = def.OK

	recvMsg := &def.GetPopularA{}
	logger.Info("recv_data = %s", string(recv_data))
	err := json.Unmarshal(recv_data, recvMsg)
	if err != nil {
		logger.Error("GetPopular json unmarshal err = %+v", err)
		send_msg.Ret = def.Ret_UnMarshalERR
		return
	}

	send_msg.Crafts = module.RankApi.GetRandomPopular()
}

func SyncInfo(recv_data []byte, send_msg *def.SyncInfoR, accid string) {
	send_msg.Ret = def.OK

	recvMsg := &def.SyncInfoA{}
	logger.Info("recv_data = %s", string(recv_data))
	err := json.Unmarshal(recv_data, recvMsg)
	if err != nil {
		logger.Error("SyncInfo json unmarshal err = %+v", err)
		send_msg.Ret = def.Ret_UnMarshalERR
		return
	}

	info, err := json.Marshal(recvMsg.Info)
	if err != nil {
		logger.Error("SyncInfo info marshal err = %+v", err)
	}
	send_msg.Ret = module.UserApi.SyncInfo(accid, string(info))
	return
}
