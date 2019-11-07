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
var produce_wait chan struct{}
var consume_wait chan struct{}

var stopProduce = false
var stopConsume = false

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
		if stopProduce {
			fmt.Println("Producer ", index, " stop produce, sleep 5 seconds")
			lock.Unlock()
			time.Sleep(time.Second * 5)
			continue
		}
		fmt.Println("Producer ", index, " begin produce")
		if productcount >= PRODUCT_MAX {
			fmt.Println("Products are full")
			stopProduce = true
			lock.Unlock()
			//产品满了，生产者wait
			<-produce_wait
			lock.Lock()
			stopProduce = false
			lock.Unlock()
			continue
		}
		productcount++
		fmt.Println("Products count is ", productcount)
		if stopConsume {
			var consumActive struct{}
			consume_wait <- consumActive
		}
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
		if stopConsume {
			fmt.Println("Consumer ", index, " stop consume, sleep 5 seconds")
			lock.Unlock()
			time.Sleep(time.Second * 5)
			continue
		}
		fmt.Println("Consumer ", index, " begin consume")
		if productcount <= 0 {
			fmt.Println("Products are empty")
			stopConsume = true
			lock.Unlock()
			//产品空了，消费者等待
			<-consume_wait
			lock.Lock()
			stopConsume = false
			lock.Unlock()
			continue
		}
		productcount--
		fmt.Println("Products count is ", productcount)

		if stopProduce {
			var productActive struct{}
			produce_wait <- productActive
		}

		lock.Unlock()
	}
}

func main() {
	wgrp.Add(PRODUCER_MAX + CONSUMER_MAX)
	produce_wait = make(chan struct{})
	consume_wait = make(chan struct{})
	for i := 0; i < CONSUMER_MAX; i++ {
		go Consume(i, &wgrp)
	}
	for i := 0; i < PRODUCER_MAX; i++ {
		go Produce(i, &wgrp)
	}

	wgrp.Wait()
}
