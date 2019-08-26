package log

import (

	// 导入生成的 protobuf 代码

	"fmt"
	config "golang-/grpcservice/serviceconfig"
	"log"
	"os"
	"sync"
)

type LogManager struct {
	File *os.File
	*log.Logger
}

func NewLogManager() (*LogManager, error) {
	file, err := os.OpenFile(logpath, os.O_APPEND|os.O_CREATE, 666)
	if err != nil {
		return nil, config.ErrLogInit
	}

	logger := log.New(file, "", log.LstdFlags|log.Lshortfile) // 日志文件格式:log包含时间及文件行数

	lm := new(LogManager)
	lm.File = file
	lm.Logger = logger

	//log.Println("输出日志到命令行终端")
	//logger.Println("将日志写入文件")
	return lm, nil
}

func InitLog(path string) *LogManager {
	if path != "" {
		logpath = path
	}
	return GetLogManagerIns()
}

func (lm *LogManager) CloseLogMgr() {
	if lm.File == nil {
		return
	}
	lm.File.Close()
	if lm.Logger == nil {
		return
	}

}

var logmgr *LogManager = nil
var once sync.Once
var logpath string = "./logdata.log"

func GetLogManagerIns() *LogManager {
	once.Do(func() {
		var err error
		logmgr, err = NewLogManager()
		fmt.Printf("Log init , path is %s", logpath)
		fmt.Println("")
		if err != nil {
			fmt.Println("Log init failed ", err.Error())
		}
	})
	return logmgr
}
