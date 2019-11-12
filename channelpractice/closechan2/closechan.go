package main

import "fmt"

//可以执行go run -race 进行死锁检测，看到主线程挂起
func main() {
	nochan := make(chan int)
	go func(ch chan int) {
		close(ch)
		fmt.Println("goroutine exit")
	}(nochan)

	data, ok := <-nochan
	if !ok {
		fmt.Println("receive close chan")
		fmt.Println("receive data is ", data)
	}
	//二次关闭
	close(nochan)
	fmt.Println("main exited")
}
