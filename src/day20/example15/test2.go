package main

import "fmt"

func Adder() func(int) int {
	var x int
	return func(delta int) int {
		fmt.Println("x is : ", x)
		x += delta
		return x
	}
}

func main() {
	var f func(int) int = Adder()
	fmt.Println("f(1)", f(1))
	fmt.Println("f(20) ", f(20))
	fmt.Println("f(300)", f(300))
}
