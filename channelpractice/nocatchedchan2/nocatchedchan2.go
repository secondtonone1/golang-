package main

import (
	"fmt"
	"sync"
)

func main() {
	nochan := make(chan int)
	waiter := &sync.WaitGroup{}
	waiter.Add(2)
	go func(ch chan int, wt *sync.WaitGroup) {
		data := <-ch
		fmt.Println("receive data ", data)
		wt.Done()
	}(nochan, waiter)

	go func(ch chan int, wt *sync.WaitGroup) {
		ch <- 5
		fmt.Println("send data ", 5)
		wt.Done()
	}(nochan, waiter)
	waiter.Wait()
}
