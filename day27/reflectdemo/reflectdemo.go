package main

import (
	"fmt"
	"reflect"
)

type Hero struct {
	name string
	id   int
}

func (h Hero) PrintData() {
	fmt.Println("Hero name is ", h.name, " id is ", h.id)
}

func (h Hero) SetName(name string) {
	h.name = name
}

func (h *Hero) SetName2(name string) {
	h.name = name
}

func (h *Hero) PrintData2() {
	fmt.Println("Hero name is ", h.name, " id is ", h.id)
}

func ReflectTypeValue(itf interface{}) {

	rtype := reflect.TypeOf(itf)
	fmt.Println("reflect type is ", rtype)
	rvalue := reflect.ValueOf(itf)
	fmt.Println("reflect value is ", rvalue)
	fmt.Println("reflect  value kind is", rvalue.Kind())
	fmt.Println("reflect type kind is", rtype.Kind())
	fmt.Println("reflect  value type is", rvalue.Type())
}

func ReflectStructElem(itf interface{}) {
	rvalue := reflect.ValueOf(itf)
	for i := 0; i < rvalue.NumField(); i++ {
		elevalue := rvalue.Field(i)
		fmt.Println("element ", i, " its type is ", elevalue.Type())
		fmt.Println("element ", i, " its kind is ", elevalue.Kind())
		fmt.Println("element ", i, " its value is ", elevalue)
	}
}

func ReflectStructPtrElem(itf interface{}) {
	rvalue := reflect.ValueOf(itf)
	for i := 0; i < rvalue.Elem().NumField(); i++ {
		elevalue := rvalue.Elem().Field(i)
		fmt.Println("element ", i, " its type is ", elevalue.Type())
		fmt.Println("element ", i, " its kind is ", elevalue.Kind())
		fmt.Println("element ", i, " its value is ", elevalue)
	}

	if rvalue.Elem().Field(1).CanSet() {
		rvalue.Elem().Field(1).SetInt(100)
	} else {
		fmt.Println("struct element 1 can't be changed")
	}

}

func ReflectStructMethod(itf interface{}) {
	rvalue := reflect.ValueOf(itf)
	rtype := reflect.TypeOf(itf)
	for i := 0; i < rvalue.NumMethod(); i++ {
		methodvalue := rvalue.Method(i)
		fmt.Println("method ", i, " value is ", methodvalue)
		methodtype := rtype.Method(i)
		fmt.Println("method ", i, " type is ", methodtype)
		fmt.Println("method ", i, " name is ", methodtype.Name)
		fmt.Println("method ", i, " method.type is ", methodtype.Type)
	}

	//reflect.ValueOf 方法调用,无参方法调用
	fmt.Println(rvalue.Method(0).Call(nil))
	//有参方法调用
	params := []reflect.Value{reflect.ValueOf("Rolin")}
	fmt.Println(rvalue.Method(1).Call(params))
	//虽然修改了，但是并没有生效
	fmt.Println(rvalue.Method(0).Call(nil))
}

func ReflectStructPtrMethod(itf interface{}) {
	rvalue := reflect.ValueOf(itf)
	rtype := reflect.TypeOf(itf)
	for i := 0; i < rvalue.NumMethod(); i++ {
		methodvalue := rvalue.Method(i)
		fmt.Println("method ", i, " value is ", methodvalue)
		methodtype := rtype.Method(i)
		fmt.Println("method ", i, " type is ", methodtype)
		fmt.Println("method ", i, " name is ", methodtype.Name)
		fmt.Println("method ", i, " method.type is ", methodtype.Type)
	}

	//reflect.ValueOf 方法调用,无参方法调用
	fmt.Println(rvalue.Method(1).Call(nil))
	//有参方法调用
	params := []reflect.Value{reflect.ValueOf("Rolin")}
	fmt.Println(rvalue.Method(3).Call(params))
	//修改了，生效
	fmt.Println(rvalue.Method(0).Call(nil))

	for i := 0; i < rvalue.Elem().NumMethod(); i++ {
		methodvalue := rvalue.Elem().Method(i)
		fmt.Println("method ", i, " value is ", methodvalue)
		methodtype := rtype.Elem().Method(i)
		fmt.Println("method ", i, " type is ", methodtype)
		fmt.Println("method ", i, " name is ", methodtype.Name)
		fmt.Println("method ", i, " method.type is ", methodtype.Type)
	}
}

func GetMethodByName(itf interface{}) {
	rvalue := reflect.ValueOf(itf)
	methodvalue := rvalue.MethodByName("PrintData2")
	if !methodvalue.IsValid() {
		return
	}

	methodvalue.Call(nil)

	methodset := rvalue.MethodByName("SetName2")
	if !methodset.IsValid() {
		return
	}
	params := []reflect.Value{reflect.ValueOf("Hurricane")}
	methodset.Call(params)

	methodvalue.Call(nil)
}

func main() {
	var num float64 = 13.14
	rtype := reflect.TypeOf(num)
	fmt.Println("reflect type is ", rtype)
	rvalue := reflect.ValueOf(num)
	fmt.Println("reflect value is ", rvalue)
	fmt.Println("reflect  value kind is", rvalue.Kind())
	fmt.Println("reflect type kind is", rtype.Kind())
	fmt.Println("reflect  value type is", rvalue.Type())

	rptrvalue := reflect.ValueOf(&num)
	rptrvalue.Elem().SetFloat(131.4)
	fmt.Println(num)
	//rvalue 为reflect包的Value类型
	//可通过Interface()转化为interface{}类型，进而转化为原始类型
	rawvalue := rvalue.Interface().(float64)
	fmt.Println("rawvalue is ", rawvalue)
	ReflectTypeValue(Hero{name: "zack", id: 1})
	ReflectTypeValue(&Hero{name: "zack", id: 1})
	ReflectStructElem(Hero{name: "zack fair", id: 2})
	//ReflectStructElem(&Hero{name: "Rolin", id: 20})
	heroptr := &Hero{name: "zack fair", id: 2}
	ReflectStructPtrElem(heroptr)
	ReflectStructMethod(Hero{name: "zack fair", id: 2})

	ReflectStructPtrMethod(&Hero{name: "zack fair", id: 2})

	GetMethodByName(&Hero{name: "zack fair", id: 2})
}
