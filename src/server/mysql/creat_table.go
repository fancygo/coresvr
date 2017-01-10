package main

//数据库处理, 直接使用gorm库, 包括建库等

import (
	_ "encoding/binary"
	"fmt"
	"frame"
	"frame/def"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "io/ioutil"
	_ "math"
	_ "math/rand"
	_ "runtime"
	_ "time"
)

var gDB *gorm.DB

func init() {
	var err error
	gDB, err = connectMysql()
	if err != nil {
		panic(fmt.Sprintf("Connect mysql error = %v\n", err))
	}
	fmt.Println("Connect success")
	gDB.LogMode(false)
	gDB.DB().SetMaxIdleConns(10)
}

func connectMysql() (db *gorm.DB, err error) {
	sqlId := frame.GetMysqlId()
	user := frame.GetSqlUser(sqlId)
	pswd := frame.GetSqlPswd(sqlId)
	host := frame.GetSqlHost(sqlId)
	port := frame.GetSqlPort(sqlId)
	database := "game"
	login := user + ":" + pswd + "@(" + host + ":" + port + ")/" + database + "?charset=utf8&parseTime=True&loc=Local"
	fmt.Println(login)
	db, err = gorm.Open("mysql", login)
	return
}

func initTable() {
	if ok := gDB.HasTable(&def.User{}); !ok {
		//gDB.DropTable(&def.User{})
		gDB.CreateTable(&def.User{})
	}
}

func initData() {
	for i := 0; i < 10; i++ {
		user := def.User{
			Gold: 0,
			Info: "",
		}
		gDB.Create(&user)
	}
}

func main() {
	fmt.Println("this mysql tool")
	initTable()
	//initData()
}
