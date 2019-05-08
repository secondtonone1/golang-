package main

import (
	"fmt"
	"time"
)

//单项读
func ReadFromChan(readchan <-chan int) {
	data := <-readchan
	fmt.Println(data)
}

//单项写
func WriteToChan(writechan chan<- int) {
	data := 121
	writechan <- data
}

// 判断是否关闭chan
func JudgeClose(ch <-chan int) {
	data, res := <-ch
	if !res {
		fmt.Println("chan has been closed")
		return
	}
	fmt.Println(data)

}

func main() {
	ch := make(chan int, 1)
	ch <- 1
	go ReadFromChan(ch)
	time.Sleep(time.Duration(1) * time.Second)
	go WriteToChan(ch)
	time.Sleep(time.Duration(1) * time.Second)
	data := <-ch
	fmt.Println(data)
	//关闭ch
	close(ch)
	go JudgeClose(ch)
	time.Sleep(time.Duration(1) * time.Second)
}
