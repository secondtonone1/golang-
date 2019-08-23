package log

import (

	// 导入生成的 protobuf 代码

	config "golang-/grpcservice/serviceconfig"
	"log"
	"os"
)

type LogManager struct {
	File   *os.File
	Logger *log.Logger
}

func NewLogManager(path string) (*LogManager, error) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE, 666)
	if err != nil {
		return nil, config.ErrLogInit
	}

	logger = log.New(file, "", log.LstdFlags|log.Lshortfile) // 日志文件格式:log包含时间及文件行数

	lm := new(LogManager)
	lm.File = file
	lm.Logger = logger

	defer file.Close()

	log.Println("输出日志到命令行终端")
	logger.Println("将日志写入文件")
}
