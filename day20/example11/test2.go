package main

import (
	"fmt"
	"strconv")

func main() {
	str:=""
	fmt.Scanf("%s", &str)
	num,err:=strconv.Atoi(str)
	if err != nil{
		fmt.Println("input error ")
	}
	switch num {
	case 1:
		fmt.Println("num is 1")
		fallthrough
	case 2:
		fmt.Println("num is <= 2")
	case 3:
		fmt.Println("num is 3")
	}
	
	switch num{
	case 1,2:
		fmt.Println("num is <= 2")
	case 3:
		fmt.Println("num is 3")
	default:
		fmt.Println("default logic")
	}

	switch {
	case num <=2 && num >= 0:
		fmt.Println("num is between 0 and 2")
	case num >=3 && num <=5:
		fmt.Println("num is between 3 and 5")
	default:
		fmt.Println("default logic")
	}
}
