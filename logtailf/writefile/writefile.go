package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"runtime"
	"time"
)

func writeLog(datapath string) {
	filew, err := os.OpenFile(datapath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("open file error ", err.Error())
		return
	}

	w := bufio.NewWriter(filew)
	for i := 0; i < 20; i++ {
		timeStr := time.Now().Format("2006-01-02 15:04:05")
		fmt.Fprintln(w, "Hello current time is "+timeStr)
		time.Sleep(time.Millisecond * 100)
		w.Flush()
	}
	logBak := time.Now().Format("20060102150405") + ".txt"
	logBak = path.Join(path.Dir(datapath), logBak)
	filew.Close()
	err = os.Rename(datapath, logBak)
	if err != nil {
		fmt.Println("Rename error ", err.Error())
		return
	}
}

func main() {
	logrelative := `../logdir1/log.txt`
	_, filename, _, _ := runtime.Caller(0)
	fmt.Println(filename)
	datapath := path.Join(path.Dir(filename), logrelative)
	// fmt.Println(datapath)
	// fmt.Println(path.Base(datapath))
	// fmt.Println(strings.Split(path.Base(datapath), ".")[1])
	// fmt.Println(path.Dir(datapath))
	for i := 0; i < 3; i++ {
		writeLog(datapath)
	}
}
