package main

import "fmt"

//可以执行go run -race 进行死锁检测，看到主线程挂起
func main() {
	var nochan chan int = nil
	go func(ch chan int) {
		//关闭nil channel会panic
		close(ch)
		fmt.Println("goroutine exit")
	}(nochan)

	//从nil channel中读取会阻塞
	data, ok := <-nochan
	if !ok {
		fmt.Println("receive close chan")
		fmt.Println("receive data is ", data)
	}
	fmt.Println("main exited")
}
