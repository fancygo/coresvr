package frame

//文件目录处理

import (
	"os"
	_ "os/exec"
	"path"
)

var (
	bInit   = false
	workDir string
	confDir string
	binDir  string
	srcDir  string
	pidDir  string
	logDir  string
)

func GetWorkDir() string {
	return workDir
}

func GetBinDir() string {
	return binDir
}

func GetConfDir() string {
	return confDir
}

func GetLogDir() string {
	return logDir
}

func GetPidDir() string {
	return pidDir
}

func GetSrcDir() string {
	return srcDir
}

func init() {
	if bInit == true {
		return
	}
	dirInit("")
	bInit = true
	LoadConfig()
}

func dirInit(wd string) {
	if wd == "" {
		workDir, _ = os.Getwd()
		binDir = path.Join(workDir, "bin")
		confDir = path.Join(workDir, "conf")
		pidDir = path.Join(workDir, "pid")
		logDir = path.Join(workDir, "log")
		srcDir = path.Join(workDir, "src")
	}
}
