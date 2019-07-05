package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Structjson struct {
	UserId int    `json:"user_id"`
	Name   string `json:"user_name"`
}

func main() {
	var structp *Structjson = &Structjson{1, "John"}
	jsons, err_ := json.Marshal(structp)
	if err_ != nil {
		fmt.Println("json Marshal error")
		return
	}

	fmt.Println(string(jsons))
	//通过反射获取
	field := reflect.TypeOf(structp).Elem().Field(0)
	fmt.Println(field.Tag.Get("json"))

}
