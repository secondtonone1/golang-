package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		a := rand.Intn(100)
		println(a)
	}

	for i := 0; i < 10; i++ {
		a := rand.Float32()
		print(a)
	}
	fmt.Println("test1 called Call")
}
