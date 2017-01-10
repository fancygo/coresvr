package logger

//自己实现的log处理

import (
	"fmt"
	"log"
	"os"
)

var gLogger *log.Logger
var gStdLogger *log.Logger
var console bool

func Init(flag bool, svr_name string) {
	console = flag
	colorInit()
	logFile := "log/" + svr_name + "_server.log"
	file, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Fatalln("fail to create game.log")
	}
	gLogger = log.New(file, "", log.LstdFlags|log.Lshortfile)
	gStdLogger = log.New(os.Stderr, "", log.Lshortfile|log.LstdFlags|log.Lmicroseconds)
}

func SetConsole(flag bool) {
	console = flag
}

func output(str string, colorStr string) {
	if console == true {
		gStdLogger.Output(2, colorStr)
	} else {
		gLogger.Output(2, str)
	}
}

func Trace(s string, v ...interface{}) {
	str := fmt.Sprintf(s, v...)
	str = "[TRACE]" + str
	colorStr := colorString(COLOR_1, str)
	output(str, colorStr)
}

func Debug(s string, v ...interface{}) {
	str := fmt.Sprintf(s, v...)
	str = "[DEBUG]" + str
	colorStr := colorString(COLOR_2, str)
	output(str, colorStr)
}

func Info(s string, v ...interface{}) {
	str := fmt.Sprintf(s, v...)
	str = "[INFO]" + str
	colorStr := colorString(COLOR_3, str)
	output(str, colorStr)
}

func Warn(s string, v ...interface{}) {
	str := fmt.Sprintf(s, v...)
	str = "[WARN]" + str
	colorStr := colorString(COLOR_5, str)
	output(str, colorStr)
}

func Error(s string, v ...interface{}) {
	str := fmt.Sprintf(s, v...)
	str = "[ERROR]" + str
	colorStr := colorString(COLOR_6, str)
	output(str, colorStr)
}

func Fatal(s string, v ...interface{}) {
	str := fmt.Sprintf(s, v...)
	str = "[FATAL]" + str
	colorStr := colorString(COLOR_7, str)
	output(str, colorStr)
}
