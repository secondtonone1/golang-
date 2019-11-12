package main

import "fmt"

//可以执行go run -race 进行死锁检测，看到主线程挂起
func main() {
	nochan := make(chan int)
	go func(ch chan int) {
		ch <- 100
		fmt.Println("send data", 100)
		close(ch)
		fmt.Println("goroutine exit")
	}(nochan)

	data := <-nochan
	fmt.Println("receive data is ", data)
	//从关闭的
	data, ok := <-nochan
	if !ok {
		fmt.Println("receive close chan")
		fmt.Println("receive data is ", data)
	}
	fmt.Println("main exited")
}
