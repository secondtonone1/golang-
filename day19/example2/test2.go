package main

import (
	"fmt"
	"math/rand"
	"time"
)

func rand_generator(n int) chan int {
	rand.Seed(time.Now().UnixNano())
	out := make(chan int)
	go func(a int) {
		for {
			out <- rand.Intn(a)
		}
	}(n)

	return out
}

func main() {

	var out chan int = rand_generator(100)
	for {
		a := <-out
		fmt.Printf("%d\n", a)
		time.Sleep(time.Duration(1) * time.Second)
	}
}
