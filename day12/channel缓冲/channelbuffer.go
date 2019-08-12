package main

import (
	"fmt"
	"time"
)

func ReadFromChan(cp chan int) {
	//遍历读取chan中数据
	fmt.Println("111111111")
	for i := range cp {
		fmt.Println(i)
	}
	fmt.Println("222222222")
}

func main() {
	//创建一个带缓冲大小的chan

	c := make(chan int, 1024)
	//可以循环写入，知道缓冲区填满
	for i := 0; i < 10; i++ {
		c <- i
	}

	go ReadFromChan(c)
	time.Sleep(time.Duration(2) * time.Second)
}
