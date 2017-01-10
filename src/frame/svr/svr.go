package svr

import (
	"flag"
	"frame"
	"frame/logger"
)

type Svr struct {
	Id          int
	Name        string
	FlagConsole *bool
}

func NewSvr(id int) *Svr {
	return &Svr{
		Id: id,
	}
}

func (this *Svr) Init() {
	this.Name = frame.GetSvrName(this.Id)
	this.regFlags()
	flag.Parse()
	this.initLogger()
}

func (this *Svr) regFlags() {
	this.FlagConsole = flag.Bool("c", false, "show log in console")
}

func (this *Svr) initLogger() {
	logger.SetConsole(*this.FlagConsole)
}
