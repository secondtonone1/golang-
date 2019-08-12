package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var mutexl sync.Mutex
var rwmutexl sync.RWMutex
var count int32

func testLock(data map[int]int) {
	for i := 0; i < 2; i++ {
		go func() {
			mutexl.Lock()
			data[rand.Intn(100)] = rand.Intn(100)
			mutexl.Unlock()
		}()
	}

	for i := 0; i < 100; i++ {
		fmt.Println(i)
		go func() {
			for {
				mutexl.Lock()
				time.Sleep(time.Millisecond * time.Duration(1))
				/*
					for _, v := range data {
						fmt.Println(v)
					}*/
				mutexl.Unlock()
				atomic.AddInt32(&count, 1)
			}

		}()
	}
	time.Sleep(time.Second * time.Duration(2))
	fmt.Println(atomic.LoadInt32(&count))
}

func testRWLock(data map[int]int) {
	for i := 0; i < 2; i++ {
		go func() {
			rwmutexl.Lock()
			data[rand.Intn(100)] = rand.Intn(100)
			rwmutexl.Unlock()
		}()
	}

	for i := 0; i < 100; i++ {
		fmt.Println(i)
		go func() {
			for {
				rwmutexl.RLock()
				time.Sleep(time.Millisecond * time.Duration(1))
				/*
					for _, v := range data {
						fmt.Println(v)
					}
				*/
				rwmutexl.RUnlock()
				atomic.AddInt32(&count, 1)
			}

		}()
	}
	time.Sleep(time.Second * time.Duration(2))
	fmt.Println(atomic.LoadInt32(&count))
}

//go build -race 选项后，会自动检测锁
func main() {
	data := make(map[int]int)
	rand.Seed(time.Now().UnixNano())
	testLock(data)
	//testRWLock(data)
}
