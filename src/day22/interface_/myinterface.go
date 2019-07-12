package main

import "fmt"

//Fish 是一个接口
// 接口声明了一个Swim方法
type Fish interface {
	Swim() string
}

//定义结构体Tiddler
type Tiddler struct {
	sw string
}

//实现swim
//理论上所有实现Swim方法的结构体都可以赋值给Fish
func (td *Tiddler) Swim() string {
	fmt.Println("Tiddler can swim ", td.sw)
	return td.sw
}

func SwimFunc(fs Fish) {
	fs.Swim()
}

//接口也可以继承
//Interface类型可以定义一组方法，但是这些不需要实现。并且interface不能包含任何变量。

type Shark interface {
	Fish
	Attack()
}

func AttackFunc(sk Shark) {
	sk.Attack()
	vartype, err := sk.(*TigerShark)
	if !err {
		fmt.Println("sk is not *TigerShark type")
	} else {
		fmt.Println("sk is type of ", vartype)
	}

}

type TigerShark struct {
}

func (ts *TigerShark) Attack() {
	fmt.Println("TigerShark can attack by its tooth and tail")
}

func (ts *TigerShark) Swim() string {
	str := "TigerShark can swim very quickly"
	fmt.Println(str)
	return str
}

func main() {
	tiddler := &Tiddler{"3km"}
	SwimFunc(tiddler)
	//以下用法会出错
	//因为Tiddler实现的Swim方法是基于*Tiddler的
	//所以严格说结构体的方法接收者是指针类型，就开辟指针对象。
	//接收者是结构体类型，可以使用指针或者结构体对象
	/*
		tiddler2 := Tiddler{"3km"}
		SwimFunc(tiddler2)
	*/

	tigersrk := &TigerShark{}
	AttackFunc(tigersrk)

}
