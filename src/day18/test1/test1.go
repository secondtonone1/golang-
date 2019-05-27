package test1

import (
	_ "day18/test2"
	"fmt"
)

var greetings string = "Hello Charis Wang"
var age int = 0

func init() {
	greetings = "Hello Charis Wang called init"
	age = 31
	fmt.Println("greetings is : ", greetings)
	fmt.Println("age is : ", age)
}

func Call() {
	fmt.Println("test1 called Call")
}
