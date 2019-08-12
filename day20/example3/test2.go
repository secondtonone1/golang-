package main

import "fmt"

//计算1!+2!+3!+4!+...+n!
func CalcN(num int) int {
	sum := 0
	for i := 1; i <= num; i++ {
		sum += RecursiveCalc(i)
	}

	return sum
}

func RecursiveCalc(num int) int {
	if num == 1 {
		return 1
	}

	return num * RecursiveCalc(num-1)
}

func CalcN2(num int) int {
	sum := 0
	begin := 1
	for i := 1; i <= num; i++ {
		begin = begin * i
		sum += begin
	}
	return sum
}

func main() {

	fmt.Println(CalcN(3))
	fmt.Println(CalcN(4))
	fmt.Println(CalcN2(4))
}
