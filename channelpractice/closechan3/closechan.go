package main

import "fmt"

//可以执行go run -race 进行死锁检测，看到主线程挂起
func main() {
	nochan := make(chan int)
	go func(ch chan int) {
		close(ch)
		fmt.Println("goroutine1 exit")
	}(nochan)

	data, ok := <-nochan
	if !ok {
		fmt.Println("receive close chan")
		fmt.Println("receive data is ", data)
	}

	go func(ch chan int) {
		<-ch
		fmt.Println("goroutine2 exit")
	}(nochan)

	//向关闭的channel中写数据
	nochan <- 200
	fmt.Println("main exited")
}
