package main

import (
	"fmt"
	kafconsumer "golang-/logcatchsys/kafconsumer"
	"golang-/logcatchsys/logconfig"
)

func main() {
	v := logconfig.InitVipper()
	if v == nil {
		fmt.Println("vipper init failed!")
		return
	}

	kafconsumer.GetMsgFromKafka()
}
