package dbdata

//数据库的加载存储处理, 这边需要优化, 可以改成消息队列的模式, 这样就不用加锁了

import (
	"frame"
	"frame/def"
	"frame/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "time"
)

const (
	DB_NAME = "game"
)

var (
	db *gorm.DB
)

func Init() bool {
	sqlId := frame.GetMysqlId()
	user := frame.GetSqlUser(sqlId)
	pswd := frame.GetSqlPswd(sqlId)
	host := frame.GetSqlHost(sqlId)
	dbport := frame.GetSqlPort(sqlId)
	login := user + ":" + pswd + "@(" + host + ":" + dbport + ")/" + DB_NAME + "?charset=utf8&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open("mysql", login)
	frame.CheckErr(err)
	return true
}

func LoadUser() []def.User {
	var user []def.User
	db.Find(&user)
	logger.Info("LoadUser len = %d", len(user))
	return user
}

func LoadOneUserByAccid(accid string, user *def.User) {
	db.Where(&def.User{Accid: accid}).First(user)
}

func UpdateUser(db_user *def.User) {
	db.Save(db_user)
	logger.Info("update user accidid = %d", db_user.Accid)
}

func LoadNormalCraft() []def.NormalCraft {
	var craft []def.NormalCraft
	db.Find(&craft)
	logger.Info("LoadNormalCraft len = %d", len(craft))
	return craft
}

func SaveNormalCraft(id int, author string, rect int, data string, praise int) {
	var craft def.NormalCraft
	craft.Id = id
	craft.Author = author
	craft.Rect = rect
	craft.Data = data
	craft.Praise = praise
	db.Save(&craft)
	logger.Info("save normal craft = %+v", craft)
}

func LoadGoodCraft() []def.GoodCraft {
	var craft []def.GoodCraft
	db.Find(&craft)
	logger.Info("LoadGoodCraft len = %d", len(craft))
	return craft
}

func SaveGoodCraft(id int, author string, rect int, data string, praise int) {
	var craft def.GoodCraft
	craft.Id = id
	craft.Author = author
	craft.Rect = rect
	craft.Data = data
	craft.Praise = praise
	db.Save(&craft)
	logger.Info("save good craft = %+v", craft)
}
