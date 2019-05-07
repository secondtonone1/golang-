package main

import (
	"fmt"
	"time"
)

func main() {
	//创建一个bool类型的chan变量
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(time.Duration(2) * time.Second)
		timeout <- true
	}()
	ch := make(chan int, 1)
	select {
	case <-ch:
		fmt.Println("ch selected..")
	case <-timeout:
		fmt.Println("timeout selected...")
	}

}
