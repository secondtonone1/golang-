package main

import "fmt"

//可以执行go run -race 进行死锁检测，看到主线程挂起
func main() {
	nochan := make(chan int)
	go func(ch chan int) {
		ch <- 100
		fmt.Println("send data", 100)
	}(nochan)

	for data := range nochan {
		fmt.Println("receive data is ", data)
	}

	fmt.Println("main exited")
}
