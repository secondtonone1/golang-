package main

import (
	"fmt"
	"reflect"
)

type UsrData struct {
	name  string
	age   int
	hobby string
}

func (ud *UsrData) Introduce() {
	fmt.Println("*************")
	fmt.Println("Hello everyone, my name is ", ud.name)
	fmt.Println("my age is ", ud.age)
	fmt.Println("my hobby is ", ud.hobby)
	fmt.Println("*************")
}

func CallIntroduce(inter interface{}) {
	//TypeOf.NumField
	types := reflect.TypeOf(inter)
	values := reflect.ValueOf(inter)
	for i := 0; i < types.NumField(); i++ {
		field := types.Field(i)
		fmt.Printf("field name is %s: filed type is %v\n", field.Name, field.Type)
		value := values.Field(i)
		fmt.Printf("field name is %s: filed type is %v, filed value is %v\n", field.Name, field.Type, value)
	}

}

func main() {
	//简单使用
	var f1 float32 = 1.23
	fmt.Println("type of 1.23 : ", reflect.TypeOf(f1))
	fmt.Println("value of 1.23 : ", reflect.ValueOf(f1))
	//ValueOf.interface
	//将接口类型变量转化为具体类型，类型不符合会panic
	value1 := reflect.ValueOf(&f1).Interface().(*float32)
	fmt.Println("value of &f1 is ", value1)

	value2 := reflect.ValueOf(f1).Interface().(float32)
	fmt.Println("value of &f1 is ", value2)

	//获取结构体元素
	usrdt := UsrData{"LiLei", 29, "table tennis"}
	CallIntroduce(usrdt)

}
