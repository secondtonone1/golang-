package main

import "fmt"

//defer 在被声明后，表达式中的变量值就已经确定了
func defer1(a int) int {
	defer fmt.Printf("value is %d\n", a)
	a++
	return a
}

//defer 执行顺序和栈类似，后进先出，后声明的先调用
func defer2(a int) int {
	defer fmt.Printf("value1 is %d\n", a)
	a++
	defer fmt.Printf("value2 is %d\n", a)
	return a
}

func defer3() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

func main() {
	fmt.Println("defer1 returns: ", defer1(1))
	fmt.Println("defer1 returns: ", defer2(1))
	defer3()
}
