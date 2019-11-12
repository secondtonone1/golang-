package main

import (
	"fmt"
)

func main() {
	nochan := make(chan int)

	go func(ch chan int) {
		data := <-ch
		fmt.Println("receive data ", data)
	}(nochan)

	nochan <- 5
	fmt.Println("send data ", 5)
	/*
		go func(ch chan int) {
			data := <-ch
			fmt.Println("receive data ", data)
		}(nochan)
	*/
}
