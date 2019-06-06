package main

import "fmt"

func addfunc(a int, args ...int) int {
	sum := 0
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	sum += a
	return sum
}

func concatfunc(firststr string, args ...string) string {
	str := ""
	str = str + " "+ firststr
	for _, temp := range args {
		str = str + " " +temp
	}
	return str
}

func main() {
	fmt.Println("addfunc(1, 2, 3, 4)", addfunc(1, 2, 3, 4))
	fmt.Println("addfunc(2)", addfunc(2))
	fmt.Println("concat(\"hello\",\"world\")", concatfunc("hello", "world"))
	fmt.Println("concat(\"nice\")", concatfunc("nice"))
}
