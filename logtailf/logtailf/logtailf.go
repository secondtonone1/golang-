package main

import (
	"fmt"
	"path"
	"runtime"
	"time"

	"github.com/hpcloud/tail"
)

func main() {
	logrelative := `../logdir/log.txt`
	_, filename, _, _ := runtime.Caller(0)
	fmt.Println(filename)
	datapath := path.Join(path.Dir(filename), logrelative)
	fmt.Println(datapath)
	tailFile, err := tail.TailFile(datapath, tail.Config{
		//文件被移除或被打包，需要重新打开
		ReOpen: true,
		//实时跟踪
		Follow: true,
		//如果程序出现异常，保存上次读取的位置，避免重新读取
		Location: &tail.SeekInfo{Offset: 0, Whence: 2},
		//支持文件不存在
		MustExist: false,
		Poll:      true,
	})

	if err != nil {
		fmt.Println("tail file err:", err)
		return
	}

	for true {
		msg, ok := <-tailFile.Lines
		if !ok {
			fmt.Printf("tail file close reopen, filename: %s\n", tailFile.Filename)
			time.Sleep(100 * time.Millisecond)
			continue
		}
		//fmt.Println("msg:", msg)
		//只打印text
		fmt.Println("msg:", msg.Text)
	}
}
