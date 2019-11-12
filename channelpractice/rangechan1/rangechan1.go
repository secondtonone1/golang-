package main

import "fmt"

func main() {
	catchan := make(chan int, 2)

	go func(ch chan int) {
		for i := 0; i < 2; i++ {
			ch <- i
			fmt.Println("send data is ", i)
		}
		//不关闭close，主协程将无法range退出
		close(ch)
		fmt.Println("goroutine1 exited")
	}(catchan)

	for data := range catchan {
		fmt.Println("receive data is ", data)
	}

	fmt.Println("main exited")
}
