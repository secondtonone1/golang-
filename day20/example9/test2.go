package main

import (
	"fmt"
	"strconv")

func main() {
	fmt.Println("input a num please")
	var strinput string
	fmt.Scanf("%s\n",&strinput)
	num, err:= strconv.Atoi(strinput)
	if err == nil{
		fmt.Printf("num is %d\n", num)
	}else{
		fmt.Printf("string can't convert to num")
	}

}
