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
type Shark interface {
	Fish
	Attack()
}

func AttackFunc(sk Shark) {
	sk.Attack()
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
