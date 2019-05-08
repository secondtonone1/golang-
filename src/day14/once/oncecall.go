package main

import (
	"fmt"
	"sync"
	"time"
)

var once sync.Once
var i int = 0

func AddNum() {
	i++
	fmt.Println("i is: ", i)
}

func main() {
	go once.Do(AddNum)
	go once.Do(AddNum)
	once.Do(AddNum)
	time.Sleep(time.Duration(2) * time.Second)
}
