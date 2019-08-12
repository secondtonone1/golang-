package main

import (
	"fmt"
)

func Count(ch chan int) {
	ch <- 1
	fmt.Println("Counting...")
}

func main() {
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count(chs[i])
	}

	for _, ch := range chs {
		<-ch
		fmt.Println(ch)
	}
}

/*
func main() {
	go func() { //让子协程先执行
		for i := 0; i < 5; i++ {
			runtime.Gosched()
			fmt.Println("go")
		}
	}()

	for i := 0; i < 2; i++ {
		//让出时间片，先让别的协议执行，它执行完，再回来执行此协程
		runtime.Gosched()
		fmt.Println("hello")
	}
}*/
