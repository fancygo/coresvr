package module

import (
	"encoding/json"
	"frame/def"
	"frame/logger"
	"server/game_server/dbdata"
	"sync"
)

type User struct {
	Accid string
	Name  string
	Info  string
}

func NewUser(accid string, name string) *User {
	return &User{
		Accid: accid,
		Name:  name,
		Info:  "",
	}
}

/////////////////////////////////////user mgr///////////////////////
type UserMgr struct {
	Users map[string]*User

	lock *sync.RWMutex
}

func NewUserMgr() *UserMgr {
	return &UserMgr{
		Users: make(map[string]*User),

		lock: new(sync.RWMutex),
	}
}

func (this *UserMgr) InitNewUser(accid string, name string) *User {
	//根据accid查找是否有账号,没有的话创建
	user := NewUser(accid, name)
	dbUser := &def.User{}
	dbdata.LoadOneUserByAccid(accid, dbUser)
	info := &def.UserInfo{}
	infoData, _ := json.Marshal(info)
	if len(dbUser.Accid) == 0 {
		dbUser.Accid = accid
		dbUser.Name = name
		//初始化UserInfo
		dbUser.Info = string(infoData)
		dbdata.UpdateUser(dbUser)
	}
	user.Accid = dbUser.Accid
	user.Name = dbUser.Name
	user.Info = dbUser.Info
	return user
}

func (this *UserMgr) AddUser(user *User) {
	this.Users[user.Accid] = user
	logger.Info("UserMgr AddUser accid = %s", user.Accid)
}

func (this *UserMgr) GetUser(accid string) *User {
	if user, ok := this.Users[accid]; !ok {
		return nil
	} else {
		return user
	}
}

func (this *UserMgr) DelUser(accid string) {
	delete(this.Users, accid)
	logger.Info("UserMgr DelUser accid = %s", accid)
}

func (this *UserMgr) DumpUser() {
	for _, v := range this.Users {
		logger.Info("user accid = %d", v.Accid)
	}
}

func (this *UserMgr) SyncInfo(accid string, info string) int {
	user, ok := this.Users[accid]
	if !ok {
		logger.Error("User accid = %s not login", accid)
		return def.Ret_UserNotLogin
	}
	user.Info = info
	dbUser := &def.User{
		Accid: accid,
		Name:  user.Name,
		Info:  user.Info,
	}
	dbdata.UpdateUser(dbUser)
	return def.OK
}
