package main

import "fmt"

func main() {
	catchan := make(chan int, 2)

	go func(ch chan int) {
		for i := 0; i < 2; i++ {
			ch <- i
			fmt.Println("send data is ", i)
		}
	}(catchan)

	for i := 0; i < 2; i++ {
		data := <-catchan
		fmt.Println("receive data is ", data)
	}
}
