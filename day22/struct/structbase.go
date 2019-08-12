package main

type Animal struct {
	age int
}

//为Animal 实现一个方法
func (anim *Animal) Eat() {
	println("Animal can eat ..")
}

//通过匿名包含，从而可以实现继承
type People struct {
	Animal
	Education string
}

func (peop *People) Job() {
	println("I have no Job")
}

//可以重写继承的函数
type Student struct {
	People
}

func (st *Student) Job() {
	println("My Job is learn")
}

//匿名继承，类型为指针
type Worker struct {
	*People
}

func (wk *Worker) Job() {
	println("My Job is work in factory")
}

//组合，如果成员是指针类型
type Writer struct {
	base *People
}

func (wr *Writer) Job() {
	println("My Job is write a book")
}

//多重继承
type StWriter struct {
	Writer
	Student
}

//golang中x.(type)只能在switch中使用
func main() {
	people := &People{Education: "九年义务"}
	people.age = 25
	people.Eat()
	people.Job()

	LiLei := new(Student)
	LiLei.age = 18
	LiLei.Job()

	uncleWang := new(Worker)
	uncleWang.Job()

	JinYong := new(Writer)
	JinYong.base = &People{Education: "剑桥博士"}
	JinYong.Job()
	JinYong.base.Job()

}
