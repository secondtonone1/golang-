//package main
package main

import (
	"fmt"
	"sync"
	"time"
)

//golang中没有引用传递，都是值传递
//map,slice其实是内部包含了一个指向数据地址的指针
//channel内部通过共享内存方士实现，所以也包含了指向该内存的指针
var lock sync.RWMutex

func Read(mp map[string]int, index int, ch chan int) {
	fmt.Println("index is: ", index)
	lock.RLock()
	fmt.Println("index ", index, "enter in lock")
	fmt.Println(mp)
	time.Sleep(time.Duration(3) * time.Second)
	lock.RUnlock()
	fmt.Println("index ", index, "out of lock")
	<-ch
}

func Write(mp map[string]int, index int, ch chan int) {
	fmt.Println("index is: ", index)
	lock.Lock()
	fmt.Println("index ", index, "enter in lock")
	mp["index"] = index
	fmt.Println(mp)
	time.Sleep(time.Duration(3) * time.Second)
	lock.Unlock()
	fmt.Println("index ", index, "out of lock")
	<-ch
}

func main() {
	ch := make(chan int)
	mp := make(map[string]int)
	mp["index"] = 0
	go Read(mp, 1, ch)
	go Write(mp, 3, ch)
	go Read(mp, 2, ch)
	ch <- 1
	ch <- 2
	ch <- 3
}
