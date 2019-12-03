package main

import (
	"context"
	"fmt"
	"golang-/logcatchsys/logconfig"
	"golang-/logcatchsys/logtailf"
	"sync"

	"github.com/spf13/viper"
)

var mainOnce sync.Once
var configMgr map[string]*logconfig.ConfigData

func ConstructMgr(configPaths interface{}) {
	configDatas := configPaths.(map[string]interface{})
	for conkey, confval := range configDatas {
		configData := new(logconfig.ConfigData)
		configData.ConfigKey = conkey
		configData.ConfigValue = confval.(string)
		ctx, cancel := context.WithCancel(context.Background())
		configData.ConfigCancel = cancel
		configMgr[conkey] = configData
		go logtailf.WatchLogFile(configData.ConfigValue,
			ctx)
	}
}

func main() {
	v := viper.New()
	configPaths, confres := logconfig.ReadConfig(v)
	if configPaths == nil || !confres {
		fmt.Println("read config failed")
		return
	}
	configMgr = make(map[string]*logconfig.ConfigData)
	ConstructMgr(configPaths)
	ctx, cancel := context.WithCancel(context.Background())
	pathChan := make(chan interface{})
	go logconfig.WatchConfig(ctx, v, pathChan)
	defer func() {
		mainOnce.Do(func() {
			if err := recover(); err != nil {
				fmt.Println("main goroutine panic ", err) // 这里的err其实就是panic传入的内容
			}
			cancel()
		})
	}()

	for {
		select {
		case pathData, ok := <-pathChan:
			if !ok {
				return
			}
			//fmt.Println("main goroutine receive pathData")
			//fmt.Println(pathData)
			pathDataNew := pathData.(map[string]interface{})

			for oldkey, oldval := range configMgr {
				_, ok := pathDataNew[oldkey]
				if ok {
					continue
				}
				oldval.ConfigCancel()
				delete(configMgr, oldkey)
			}

			for conkey, conval := range pathDataNew {
				oldval, ok := configMgr[conkey]
				if !ok {
					configData := new(logconfig.ConfigData)
					configData.ConfigKey = conkey
					configData.ConfigValue = conval.(string)
					ctx, cancel := context.WithCancel(context.Background())
					configData.ConfigCancel = cancel
					configMgr[conkey] = configData
					fmt.Println(conval.(string))
					go logtailf.WatchLogFile(configData.ConfigValue,
						ctx)
					continue
				}

				if oldval.ConfigValue != conval.(string) {
					oldval.ConfigValue = conval.(string)
					oldval.ConfigCancel()
					ctx, cancel := context.WithCancel(context.Background())
					oldval.ConfigCancel = cancel
					go logtailf.WatchLogFile(conval.(string),
						ctx)
					continue
				}

			}
			/*
				for mgrkey, mgrval := range configMgr {
					fmt.Println(mgrkey)
					fmt.Println(mgrval)
				}*/
		}
	}
}
