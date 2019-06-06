package main

import (
	"fmt"
	"strconv"
	"time"
	"math/rand"
)

func judgeRes(num int, res int)bool{
	var flag bool = false
	switch{
	case num < res:
		fmt.Println("input num is smaller.")
		flag = false
	case num > res:
		fmt.Println("input num is bigger.")
		flag = false
	case num == res:
		fmt.Println("yes , u guess right.")
		flag = true
	}

	return flag
}


func main() {
	strnum := ""
	fmt.Println("input a num")
	fmt.Scanf("%s\n",&strnum)
	num, err :=strconv.Atoi(strnum)
	if err != nil{
		fmt.Println("*****************")
		fmt.Println("input error, please input a num")
		return 
	}
	
	rand.Seed(time.Now().UnixNano())
	res:=rand.Intn(100)
	for {
		var bflag bool =judgeRes(num, res)
		if bflag == true{
			break
		}
		INPUT:
		fmt.Println("input a num")
		fmt.Scanf("%s\n",&strnum)
		num, err =strconv.Atoi(strnum)
		if err != nil{
			fmt.Println("input error, please input a num")
			goto INPUT 
		}
		
	}
	
}
