package main

import (
	"fmt"
	"time"
)

//声明
//chanvar是一个ElementType类型的channel
//var chanvar chan ElementType
//intch是一个int类型的channel
//var intch chan int
//mapch的value是bool类型的channel
//var mapch map[string] chan bool

//定义
//ch := make(chan int)
//写和读都会导致chan阻塞
//写入chan会导致阻塞，读取后才解除阻塞
//ch <- value
//读取到value中，读取也会阻塞，写入后就解除阻塞
// value :=<- ch

//select 用法

func goswitch() {
	start := time.Now()
	c := make(chan interface{})
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		time.Sleep(4 * time.Second)
		close(c)
	}()

	go func() {
		time.Sleep(3 * time.Second)
		ch1 <- 3
	}()

	go func() {
		time.Sleep(3 * time.Second)
		ch2 <- 5
	}()

	fmt.Println("Blocking on read...")
	select {
	case <-c:
		fmt.Printf("Unblocked %v later.\n", time.Since(start))
	case <-ch1:
		fmt.Printf("ch1 case...")
	case <-ch2:
		fmt.Printf("ch2 case...")
	default:
		fmt.Printf("default go ...")
	}
}

func goswitch2() {
	start := time.Now()
	c := make(chan interface{})
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		time.Sleep(4 * time.Second)
		close(c)
	}()

	go func() {
		time.Sleep(3 * time.Second)
		ch1 <- 3
	}()

	go func() {
		time.Sleep(3 * time.Second)
		ch2 <- 5
	}()

	fmt.Println("Blocking on read 2...")
	select {
	case <-c:
		fmt.Printf("Unblocked %v later.\n", time.Since(start))
	case <-ch1:
		fmt.Printf("ch1 case...")
	case <-ch2:
		fmt.Printf("ch2 case...")
		//default:
		//	fmt.Printf("default go ...")
	}
}

func goswitch3() {
	start := time.Now()
	c := make(chan interface{})
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		time.Sleep(4 * time.Second)
		close(c)
	}()

	go func() {
		time.Sleep(5 * time.Second)
		ch1 <- 3
	}()

	go func() {
		time.Sleep(5 * time.Second)
		ch2 <- 5
	}()

	fmt.Println("Blocking on read 3...")
	select {
	case <-c:
		fmt.Printf("Unblocked %v later.\n", time.Since(start))
	case <-ch1:
		fmt.Printf("ch1 case...")
	case <-ch2:
		fmt.Printf("ch2 case...")
		break
		//break 打断select后边的逻辑直接退出
		fmt.Printf("after break...")
	}
}

func getChan(i int, chans []chan int) chan int {
	if i < 0 || i >= len(chans) {
		return chans[len(chans)-1]
	}
	fmt.Printf("chans[%d]\n", i)
	return chans[i]
}

func getNumber(i int, nums []int) int {
	if i < 0 || i >= len(nums) {
		return nums[len(nums)-1]
	}
	fmt.Printf("nums[%d]\n", i)
	return nums[i]
}

func goswich4() {
	var ch1 chan int
	var ch2 chan int
	var ch3 = []chan int{ch1, ch2}
	var numbers = []int{1, 2, 3, 4, 5}

	select {
	case getChan(0, ch3) <- getNumber(2, numbers):
		fmt.Println("1th case is selected.")
	case getChan(1, ch3) <- getNumber(3, numbers):
		fmt.Println("2th case is selected.")
	default:
		fmt.Println("default!.")
	}

}

func main() {
	//select选择写入1或者2
	/*
		ch := make(chan int, 1)
		for {
			select {
			case ch <- 1:
			case ch <- 2:

			}
			i := <-ch
			fmt.Println(i)
		}
	*/

	//当有多个I/O操作可以完成，select 随机选择一个执行
	//当没有任何I/O操作满足，则执行default分支
	//如果没有任何I/O满足，且没有default分支，则select阻塞，直到有I/O满足才执行
	goswitch()
	//如果注释掉default分支，则会等待3s,然后随机选择ch1或ch2分支执行
	goswitch2()
	//将ch1,ch2的func修改睡眠5s，则会阻塞一段时间，执行第一个分支
	goswitch3()
	//所有channel表达式都会被求值、所有被发送的表达式都会被求值。求值顺序：自上而下、从左到右
	goswich4()

}
