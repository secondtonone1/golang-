package main

import "fmt"

//可以执行go run -race 进行死锁检测，看到主线程挂起
func main() {
	var nochan chan int = nil
	go func(ch chan int) {
		fmt.Println("goroutine begin receive data")
		data, ok := <-nochan
		if !ok {
			fmt.Println("receive close chan")
		}
		fmt.Println("receive data is ", data)
		fmt.Println("goroutine exit")
	}(nochan)

	fmt.Println("main begin send data")
	//向nil channel中写数据会阻塞
	nochan <- 100
	fmt.Println("main exited")
}
