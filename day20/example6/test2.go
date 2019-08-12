package main

import (
	"fmt"
	"time"
)

func goroutine() {
	fmt.Println("goroutine called!")
}

func main() {
	now_t := time.Now()
	fmt.Printf("now is %s\n", now_t.Format("2006/1/2 15:04:05"))
	before := now_t.UnixNano()
	go goroutine()
	time.Sleep(time.Duration(1) * time.Second)
	after := time.Now().UnixNano()
	duration := (after - before) / 1000
	fmt.Printf("programe runs %d", duration)
}
