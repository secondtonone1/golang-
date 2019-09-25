package main

import "fmt"

//import "log"
type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}

//构造函数和初始化
type Rect struct {
	x, y          float64
	width, height float64
}

func (r *Rect) GetX() float64 {
	return r.x
}

func (r *Rect) GetY() float64 {
	return r.y
}

func (r *Rect) Area() float64 {
	return r.width * r.height
}

func (r *Rect) GetWidth() float64 {
	return r.width
}

func (r *Rect) GetHeight() float64 {
	return r.height
}

//匿名组合和派生
type Base struct {
	Name string
}

func (base *Base) Foo() {
	fmt.Println("this is Base Foo")
}

func (base *Base) Bar() {
	fmt.Println("this is Base Bar")
}

type Foo struct {
	//匿名组合
	Base
}

func (foo *Foo) Foo() {
	foo.Base.Foo()
	fmt.Println("this is Foo Foo")
}

//匿名指针组合
type DerivePoint struct {
	*Base
}

func (derivep *DerivePoint) Foo() bool {
	fmt.Println("this is DerivePoint Foo")
	return true
}

type Logger struct {
	Level int
}

//重复定义，因为匿名组合默认用类型做变量名
type MyJob struct {
	*Logger
	Name string
	//*log.Logger // duplicate field Logger
}

//同名成员覆盖，基类成员被隐藏
type X struct {
	Name string
}

func (x *X) GetName() string {
	return x.Name
}

type Y struct {
	X
	Name string
}

func (y *Y) GetName() string {
	return y.Name
}

func (y *Y) GetXName() string {
	return y.X.GetName()
}

func main() {
	var a Integer = 1
	if a.Less(2) {
		fmt.Println(a, "less 2")
	}

	rect1 := new(Rect)
	fmt.Println("rect1.Area() is ", rect1.Area())
	rect2 := &Rect{}
	fmt.Println("rect2.Area() is ", rect2.Area())
	rect3 := &Rect{0, 0, 100, 200}
	fmt.Println("rect3.Area() is ", rect3.Area())
	rect4 := &Rect{width: 100, height: 200}
	fmt.Println("rect4.Area() is ", rect4.Area())

	foo := new(Foo)
	foo.Foo()

	y := new(Y)
	y.Name = "y.Name"
	y.X.Name = "x.Name"
	name1 := y.GetXName()
	name2 := y.GetName()
	fmt.Println("name1 is ", name1, "name2 is ", name2)

	derivep := new(DerivePoint)
	derivep.Base = new(Base)
	derivep.Base.Name = "derivep.Base.Name"
	fmt.Println("derivep.Base.Name is ", derivep.Base.Name, " derivep.Name is ", derivep.Name)
	derivep.Name = "derivep.Name"
	fmt.Println("derivep.Base.Name is ", derivep.Base.Name, " derivep.Name is ", derivep.Name)
	derivep.Foo()
	derivep.Base.Foo()

	derivep2 := DerivePoint{&Base{Name: "derivepbase2"}}
	derivep2.Foo()
	fmt.Println("derivep2 base name is ", derivep2.Base.Name)
	fmt.Println("derivep2 name is ", derivep2.Name)

}
