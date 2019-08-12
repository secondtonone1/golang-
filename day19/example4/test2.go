package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	p1 := Person{name: "LiLei", age: 10}
	//字段名
	fmt.Printf("Person is : %+v\n", p1)
	//%v      相应值的默认格式
	fmt.Printf("Person is : %v\n", p1)
	//相应值的Go语法
	fmt.Printf("Person is : %#v\n", p1)
	//%T      相应值的类型的Go语法表示
	fmt.Printf("Person is : %T\n", p1)
	//%%      字面上的百分号，并非值的占位符
	fmt.Printf("Person is : %%%v\n", p1)
	var b1 bool = false
	//%t          true 或 false。
	fmt.Printf("%t\n", b1)
	//%b      二进制表示
	fmt.Printf("%b\n", 5)
	var char1 byte = 'c'
	fmt.Printf("%v\n", char1)
	fmt.Printf("%c\n", char1)
	//%d      十进制表示
	fmt.Printf("%d\n", 0x12)
	//%o      八进制表示
	fmt.Printf("%d\n", 10)
	//%x      十六进制表示，字母形式为小写 a-f
	fmt.Printf("%x\n", 13)
	//%X      十六进制表示，字母形式为大写 A-F
	fmt.Printf("%x\n", 13)
	var chararray []byte = []byte("nice to meet u")
	//%s      输出字符串表示（string类型或[]byte)
	fmt.Printf("%s\n", chararray)
	var str1 string = "and u?"
	fmt.Printf("%s\n", str1)
	str2 := fmt.Sprintf("str2 = %d\n", p1.age)
	fmt.Println(str2)
	fmt.Printf("%q\n", "nice to meet u")
	str3 := `人生如梦
	一樽还酹江月`
	fmt.Printf("%s", str3)
}
