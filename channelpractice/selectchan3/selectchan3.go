package main

import (
	"fmt"
	"sync"
)

//可以执行go run -race 进行死锁检测，看到主线程挂起
func main() {

	datachan := make(chan int)
	groutineclose := make(chan struct{})
	mainclose := make(chan struct{})
	var onceclose sync.Once
	var readclose sync.Once
	var sendclose sync.Once
	var waitgroup sync.WaitGroup
	waitgroup.Add(2)
	go func(datachan chan int, gclose chan struct{}, mclose chan struct{}, group *sync.WaitGroup) {
		defer func() {
			onceclose.Do(func() {
				close(gclose)
			})
			sendclose.Do(func() {
				close(datachan)
				fmt.Println("send goroutine closed !")
				group.Done()
			})
		}()

		for i := 0; i < 100; i++ {
			select {
			case <-gclose:
				fmt.Println("other goroutine exited")
				return
			case <-mclose:
				fmt.Println("main goroutine exited")
				return
			case datachan <- i:

			}
		}
	}(datachan, groutineclose, mainclose, &waitgroup)

	go func(datachan chan int, gclose chan struct{}, mclose chan struct{}, group *sync.WaitGroup) {
		sum := 0
		defer func() {
			if recerr := recover(); recerr != nil {
				fmt.Println("recover !")
			}

			onceclose.Do(func() {
				fmt.Println("close gclose channel")
				close(gclose)
			})
			readclose.Do(func() {
				fmt.Println("sum is ", sum)
				fmt.Println("receive goroutine closed !")
				group.Done()
			})

		}()

		for i := 0; ; i++ {
			select {
			case <-gclose:
				fmt.Println("other goroutine exited")
				return
			case <-mclose:
				fmt.Println("main goroutine exited")
				return
			case data, ok := <-datachan:
				if !ok {
					fmt.Println("receive close chan data")
					return
				}
				sum += data
				if i >= 20 {
					panic("receive goroutine test panic !!")
				}
			}
		}
	}(datachan, groutineclose, mainclose, &waitgroup)

	defer func() {
		fmt.Println("defer main close")
		close(mainclose)
	}()

	waitgroup.Wait()
	fmt.Println("main exited")

}
