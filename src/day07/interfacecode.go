package main

import "fmt"

//import "log"
type Animal interface {
	Speak() string
}

type Cat struct {
	Name string
}

func (cat *Cat) Speak() string {
	return cat.Name + " says miao miao miao..."
}

type Dog struct {
	Name string
}

func (dog *Dog) Speak() string {
	return dog.Name + " says wang wang wang ..."
}


type Bird struct{
    Name string
}

func (bird Bird) Speak() string{
    return bird.Name+" says bi bi bi ...."
}

type People interface{
    Speak() string
    Study()
}

func main(){
    var animal1 Animal = &Cat{Name:"cat"}
    fmt.Println(animal1.Speak())

    var cat2 *Cat = new(Cat)
	cat2.Name = "cat2"
	var animal21 Animal = cat2
    fmt.Println(animal21.(Animal).Speak())
    
    animals := [] Animal{&Cat{Name:"cat"},&Dog{Name:"dog"}}
    for _, animal := range animals{
        fmt.Println(animal.Speak())
    }

    // 接口对象，类型查询
    //这个if语句判断animal1接口指向的对象实例是否是*Cat类型
    if res, ok := animal1.(*Cat); ok{
        fmt.Println("animal1 type is ",res)
    }

    var animal2 Animal = Bird{Name:"bird"}

    if res, ok := animal2.(Bird); ok{
        fmt.Println("animal1 type is ",res)
    }

    //判断animal1接口指向的对象实例是否是*Dog类型
    if res, ok := animal1.(*Dog); ok{
        fmt.Println("animal1 type is ",res)
    }

    //动物接口对象不能赋值给人类接口,因为没有实现Study方法
    /*
    Animal does not implement People (missing Study method)
    */
    //var people1 People = animal1
    //接口查询，判断是否能进行接口转换
    if res, ok := animal1.(People);ok{
        fmt.Println("animal1 type is ", res)
    }

}
