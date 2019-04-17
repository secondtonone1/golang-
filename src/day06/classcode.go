package main
import "fmt"
type Integer int 
func (a Integer) Less (b Integer) bool{
    return a < b
}

//构造函数和初始化
type Rect struct{
    x,y float64
    width, height float64
}

func (r* Rect) Area() float64{
    return r.width* r.height
}

//匿名组合和派生
type Base struct{
    Name string
}

func (base* Base) Foo() {
    fmt.Println("this is Base Foo")
}

func (base* Base) Bar(){
    fmt.Println("this is Base Bar")
}

type Foo struct{
    //匿名组合
    Base
}

func (foo* Foo) Foo(){
    foo.Base.Foo()
    fmt.Println("this is Foo Foo")
}

//同名成员覆盖，基类成员被隐藏
type X struct{
    Name string
}

func (x * X) GetName() string{
    return x.Name
}

type Y struct{
    X
    Name string
}

func (y * Y) GetName() string{
    return y.Name
}

func (y * Y) GetXName() string{
    return y.X.GetName()
}

func main(){
var a Integer = 1
if a.Less(2) {
    fmt.Println(a, "less 2")
}

rect1 := new(Rect)
fmt.Println("rect1.Area() is ",rect1.Area())
rect2 := &Rect{}
fmt.Println("rect2.Area() is ",rect2.Area())
rect3 := &Rect{0,0,100,200}
fmt.Println("rect3.Area() is ",rect3.Area())
rect4 := &Rect{width:100,height:200}
fmt.Println("rect4.Area() is ",rect4.Area())

foo := new(Foo)
foo.Foo()

y := new(Y)
y.Name = "y.Name"
y.X.Name = "x.Name"
name1 := y.GetXName()
name2 := y.GetName()
fmt.Println("name1 is ",name1, "name2 is ",name2)
}