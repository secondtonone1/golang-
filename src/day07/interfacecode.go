package main
import "fmt"
//import "log"
type Animal interface{
    Speak() string
}

type Cat struct{
    Name string
}

func (cat* Cat) Speak() string {
    return cat.Name+" says miao miao miao..."
}

type Dog struct{
    Name string
}

func (dog* Dog) Speak() string {
    return dog.Name+" says wang wang wang ..."
}

func main(){
    var animal1 Animal = &Cat{Name:"cat"}
    fmt.Println(animal1.Speak())

    animals := [] Animal{&Cat{Name:"cat"},&Dog{Name:"dog"}}
    for _, animal := range animals{
        fmt.Println(animal.Speak())
    }
}

