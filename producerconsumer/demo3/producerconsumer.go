package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	PRODUCER_MAX = 5
	CONSUMER_MAX = 2
	PRODUCT_MAX  = 20
)

var productcount = 0
var lock sync.Mutex
var wgrp sync.WaitGroup

//生产者
func Produce(index int, wgrp *sync.WaitGroup) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Producer ", index, " panic")
		}
		wgrp.Done()
	}()

	for {
		time.Sleep(time.Second)
		lock.Lock()
		fmt.Println("Producer ", index, " begin produce")
		if productcount >= PRODUCT_MAX {
			fmt.Println("Products are full")
			lock.Unlock()
			return
		}
		productcount++
		fmt.Println("Products count is ", productcount)
		lock.Unlock()
	}
}

//消费者
func Consume(index int, wgrp *sync.WaitGroup) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Consumer ", index, " panic")
		}
		wgrp.Done()
	}()

	for {
		time.Sleep(time.Second)
		lock.Lock()
		fmt.Println("Consumer ", index, " begin consume")
		if productcount <= 0 {
			fmt.Println("Products are empty")
			lock.Unlock()
			return
		}
		productcount--
		fmt.Println("Products count is ", productcount)
		lock.Unlock()
	}
}

func main() {
	wgrp.Add(PRODUCER_MAX + CONSUMER_MAX)
	for i := 0; i < PRODUCER_MAX; i++ {
		go Produce(i, &wgrp)
	}

	for i := 0; i < CONSUMER_MAX; i++ {
		go Consume(i, &wgrp)
	}
	wgrp.Wait()
}
