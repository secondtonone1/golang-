package main

import "fmt"

func IsShuiXianNum(num int) bool {
	if num > 999 || num < 100 {
		return false
	}
	hund := num / 100
	ten := (num % 100) / 10
	nu := (num) % 10

	if num == hund*hund*hund+ten*ten*ten+nu*nu*nu {
		return true
	}

	return false
}

func main() {
	for i := 100; i < 1000; i++ {
		if IsShuiXianNum(i) == true {
			fmt.Println(i)
		}
	}
}
