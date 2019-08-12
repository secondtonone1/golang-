package main

import(
	"fmt"
)

func makeFib(slice []int){
	if len(slice) <= 2{
		return 
	}

	for i:=2; i < len(slice); i++{
		slice[i] = slice[i-2]+slice[i-1]
	}
}

func printFib(slice[]int){
	for i:=0; i < len(slice); i++{
		fmt.Println(slice[i])
	}
}

func main() {
	slice := make([]int, 10)
	slice[0] = 0
	slice[1] = 1
	makeFib(slice)
	printFib(slice)
}
