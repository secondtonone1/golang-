package main

import (
	"context"
	"fmt"
	"golang-/logcatchsys/logconfig"

	"github.com/spf13/viper"
)

func main() {
	v := viper.New()
	logconfig.ReadConfig(v)

	ctx, cancel := context.WithCancel(context.Background())
	pathChan := make(chan interface{})
	logconfig.WatchConfig(ctx, v, pathChan)

	select {
	case pathData := <-pathChan:
		fmt.Println(pathData)
	}

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("main goroutine panic ", err) // 这里的err其实就是panic传入的内容
		}
		cancel()
	}()
}
