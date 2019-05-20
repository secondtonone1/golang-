package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

//固定大小的结构体，就要求结构体中不能出现[]byte这样的切片成员，
//否则Size返回-1，且不能进行正常的序列化操作。

type A struct {
	one int32
	two int64
}

func main() {
	var a A
	a.one = 12
	a.two = 24
	buf := new(bytes.Buffer)
	fmt.Println("a's size is ", binary.Size(a))
	binary.Write(buf, binary.LittleEndian, a)
	fmt.Println("after write , buf is : ", buf.Bytes())
	/*
		buf := new(bytes.Buffer)
		fmt.Println("a‘s size is ", binary.Size(a))
		binary.Write(buf, binary.LittleEndian, a)
		fmt.Println("after write ，buf is:", buf.Bytes())
	*/
}
