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

//生产者
func Produce(index int, product_ch chan *int, pgroup *sync.WaitGroup) {
	defer pgroup.Done()
	for {
		time.Sleep(time.Millisecond * 100)
		select {
		case productnum := <-product_ch:
			fmt.Println("produceer num is ", index)
			if *productnum >= PRODUCT_MAX {
				fmt.Println("products are full ")
				product_ch <- productnum
				continue
			}
			*productnum++
			product_ch <- productnum
			fmt.Println("products count is ", *productnum)
		}
	}
}

//消费者
func Consume(index int, product_ch chan *int, pgroup *sync.WaitGroup) {
	defer pgroup.Done()
	for {
		time.Sleep(time.Millisecond * 100)
		select {
		case productnum := <-product_ch:
			fmt.Println("consumer num is ", index)
			if *productnum <= 0 {
				fmt.Println("products are empty ")
				product_ch <- productnum
				continue
			}
			*productnum--
			product_ch <- productnum
			fmt.Println("products count is ", *productnum)
		}
	}
}

func main() {
	product_ch := make(chan *int)
	wg := sync.WaitGroup{}
	wg.Add(CONSUMER_MAX + PRODUCER_MAX)
	var products = 0
	for i := 0; i < PRODUCER_MAX; i++ {
		go Produce(i, product_ch, &wg)
	}

	for i := 0; i < CONSUMER_MAX; i++ {
		go Consume(i, product_ch, &wg)
	}

	product_ch <- &products
	wg.Wait()
}
