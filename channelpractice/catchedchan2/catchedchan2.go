package main

import "fmt"

func main() {
	catchan := make(chan int, 2)

	for i := 0; i < 2; i++ {
		data := <-catchan
		fmt.Println("receive data is ", data)
	}
	//死锁！！！
	go func(ch chan int) {
		for i := 0; i < 2; i++ {
			ch <- i
			fmt.Println("send data is ", i)
		}
	}(catchan)
}
