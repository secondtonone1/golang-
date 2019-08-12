package test2

import "fmt"

var greetings string = "Hello Rorolin"
var age int = 0

func init() {
	greetings = "Hello Rorolin _init function called"
	fmt.Println("greetings is : ", greetings)
	age = 30
	fmt.Println("age is : ", age)
}
