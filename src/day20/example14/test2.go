package main

import "fmt"

/*
 new会分配结构空间，并初始化为清空为零，不进一步初始化
 new之后需要一个指针来指向这个结构
 make会分配结构空间及其附属空间，并完成其间的指针初始化
 make返回这个结构空间，不另外分配一个指针

*/
func main() {
	var p *[]int = new([]int)
	fmt.Println(p)
	//fmt.Println((*p)[0]) //panic: runtime error: index out of range

	s := make([]int, 10, 20)
	s[1] = 100
	fmt.Println(s[1])

}
