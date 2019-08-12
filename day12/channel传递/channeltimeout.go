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
	//range不会主动退出，直到超时,所以为了主动退出可以判断或者close等
	//close(pipchan)
	for pipData := range pipchan {
		pipData.next <- pipData.handler(pipData.value)
		//判断pipchan大小为0，则退出
		if len(pipchan) <= 0 {
			break
		}
	}
	fmt.Println("ssssssssssss")
}

func GoroutineDeal2(pipchan chan *PipeData) {
	pipData := <-pipchan
	pipData.next <- pipData.handler(pipData.value)

}

// func GoroutineDeal3(pipchan *chan *PipeData) {
// 	for pipele := range *pipchan {
// 		//判断pipchan大小为0，则退出
// 		fmt.Println(pipele.value)
// 		if len(*pipchan) <= 0 {
// 			break
// 		}
// 	}

// 	var pipDatas PipeData
// 	pipDatas.value = 250
// 	pipDatas.handler = RetValue
// 	pipDatas.next = make(chan int, 1)
// 	*pipchan <- &pipDatas
// }

func GoroutineDeal3(pipchan chan *PipeData) {
	for pipele := range pipchan {
		//判断pipchan大小为0，则退出
		fmt.Println(pipele.value)
		if len(pipchan) <= 0 {
			break
		}
	}

	var pipDatas PipeData
	pipDatas.value = 250
	pipDatas.handler = RetValue
	pipDatas.next = make(chan int, 1)
	pipchan <- &pipDatas
	fmt.Println(&pipchan)
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

	pipchan2 <- &pipData
	fmt.Println(&pipchan2)
	go GoroutineDeal3(pipchan2)
	time.Sleep(time.Duration(2) * time.Second)
	pipdata3 := <-pipchan2
	fmt.Println(pipdata3.value)

}
