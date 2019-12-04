package logtailf

import (
	"context"
	"fmt"
	"time"

	"github.com/hpcloud/tail"
)

func WatchLogFile(pathkey string, datapath string, ctx context.Context, keychan chan<- string) {
	fmt.Println("begin goroutine watch log file ", pathkey)
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

	defer func() {
		if errcover := recover(); errcover != nil {
			fmt.Println("goroutine watch ", pathkey, " panic")
			fmt.Println(errcover)
			keychan <- pathkey
		}
	}()

	//模拟崩溃
	if pathkey == "logdir3" {
		panic("test panic ")
	}

	for true {
		select {
		case msg, ok := <-tailFile.Lines:
			if !ok {
				fmt.Printf("tail file close reopen, filename: %s\n", tailFile.Filename)
				time.Sleep(100 * time.Millisecond)
				continue
			}
			//fmt.Println("msg:", msg)
			//只打印text
			fmt.Println("msg:", msg.Text)
		case <-ctx.Done():
			fmt.Println("receive main gouroutine exit msg")
			fmt.Println("watch log file ", pathkey, " goroutine exited")
			return
		}

	}
}
