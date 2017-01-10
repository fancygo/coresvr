package frame

//该文件主要是获取服务器启动的各个配置ip,port等

import (
	"encoding/json"
	"frame/logger"
	"io/ioutil"
	"os"
	"path"
)

type SvrId struct {
	Core int `json:"core"`
	Game int `json:"game"`
	Db   int `json:"db"`
	Log  int `json:"log"`
}
type SqlId struct {
	Mysql int `json:"mysql"`
	Redis int `json:"redis"`
}
type MainSvr struct {
	Id int `json:"id"`
}
type Svr struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	IP   string `json:"ip"`
	Port string `json:"port"`
}
type Sql struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Host string `json:"host"`
	Port string `json:"port"`
	User string `json:"user"`
	Pswd string `json:"pswd"`
}

type ServerCfgJson struct {
	Svrid   SvrId   `json:"svrid"`
	Sqlid   SqlId   `json:"sqlid"`
	Mainsvr MainSvr `json:"mainsvr"`
	Svr     []*Svr  `json:"svr"`
	Sql     []*Sql  `json:"sql"`
}

type SvrCfg struct {
	MainSvr int
	Sqlid   SqlId
	Svr     map[int]*Svr
	Sql     map[int]*Sql
}

var svrcfgjson ServerCfgJson
var svrcfg SvrCfg

func LoadConfig() {
	//先直接解析json文件
	configName := path.Join(GetConfDir(), "config.json")
	file, err := os.Open(configName)
	if err != nil {
		logger.Debug("error = %+v", err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		logger.Debug("error = %+v", err)
		return
	}
	err = json.Unmarshal(data, &svrcfgjson)
	if err != nil {
		logger.Debug("error = %+v", err)
		return
	}

	//将配置存放到自定义结构中
	svrcfg = SvrCfg{
		MainSvr: svrcfgjson.Mainsvr.Id,
		Sqlid: SqlId{
			Mysql: svrcfgjson.Sqlid.Mysql,
			Redis: svrcfgjson.Sqlid.Redis,
		},
		Svr: make(map[int]*Svr),
		Sql: make(map[int]*Sql),
	}
	for _, v := range svrcfgjson.Svr {
		svrcfg.Svr[v.Id] = v
	}
	for _, v := range svrcfgjson.Sql {
		svrcfg.Sql[v.Id] = v
	}
}

func GetMainSvr() int {
	return svrcfg.MainSvr
}
func GetSvrIP(id int) string {
	return svrcfg.Svr[id].IP
}
func GetSvrPort(id int) string {
	return svrcfg.Svr[id].Port
}

func GetMysqlId() int {
	return svrcfg.Sqlid.Mysql
}
func GetRedisId() int {
	return svrcfg.Sqlid.Redis
}
func GetSqlHost(id int) string {
	return svrcfg.Sql[id].Host
}
func GetSqlPort(id int) string {
	return svrcfg.Sql[id].Port
}
func GetSqlUser(id int) string {
	return svrcfg.Sql[id].User
}
func GetSqlPswd(id int) string {
	return svrcfg.Sql[id].Pswd
}
