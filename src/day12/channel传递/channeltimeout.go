package main

import (
	"fmt"
	"time"
)

type PipeData struct {
	value   int
	handler func(int) int
	next    chan int
}

func RetValue(i int) int {
	return i
}

func GoroutineDeal(pipchan chan *PipeData) {
	fmt.Println("wwwwwwwwwwwwwwwwwwww")
	for pipData := range pipchan {
		pipData.next <- pipData.handler(pipData.value)
	}
	fmt.Println("ssssssssssss")
}

func GoroutineDeal2(pipchan chan *PipeData) {
	pipData := <-pipchan
	pipData.next <- pipData.handler(pipData.value)

}

func main() {
	//创建一个PipeData类型的chan变量
	pipchan := make(chan *PipeData, 1)
	var pipData PipeData
	pipData.value = 1
	pipData.handler = RetValue
	pipData.next = make(chan int, 1)
	pipchan <- &pipData
	go GoroutineDeal(pipchan)
	time.Sleep(time.Duration(2) * time.Second)

	nextdata := <-pipData.next
	fmt.Println(nextdata)

	pipchan2 := make(chan *PipeData, 1)
	pipData.value = 200
	pipchan2 <- &pipData

	go GoroutineDeal2(pipchan2)
	time.Sleep(time.Duration(2) * time.Second)
	nextdata2 := <-pipData.next
	fmt.Println(nextdata2)

}
