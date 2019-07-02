package main

import (
	"fmt"
	"time"
)

type HeapEle struct {
	//任务的回调函数
	callback func(interface{}) bool
	//回调函数的参数
	params interface{}
	//任务设置的定时器时间
	deadline int64
}

/*
func Calback(params interface{}) bool {
	switch params.(type) {
	case string:
		fmt.Println(params)
	case []string:
		strs := params.([]string)
		for _, str := range strs {
			fmt.Println(str)
		}
	case []int:
		nums := params.([]int)
		for _, num := range nums {
			fmt.Println(num)
		}
	}
	return true
}
*/

func main() {
	//任务截止时间
	deadline := time.Now().Unix() + (int64)(10)
	heape := new(HeapEle)
	heape.deadline = deadline
	//任务参数
	heape.params = []int{1, 2, 3}
	//任务回调函数
	heape.callback = func(params interface{}) bool {
		switch params.(type) {
		case string:
			fmt.Println(params)
		case []string:
			strs := params.([]string)
			for _, str := range strs {
				fmt.Println(str)
			}
		case []int:
			nums := params.([]int)
			for _, num := range nums {
				fmt.Println(num)
			}
		}
		return true
	}
	time.Sleep(time.Duration(time.Second * 11))
	if time.Now().Unix() >= heape.deadline {
		heape.callback(heape.params)
	}

}
