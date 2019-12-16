package main

import (
	"context"
	"fmt"
	kafkaqueue "golang-/logcatchsys/kafka"
	"golang-/logcatchsys/logconfig"
	"golang-/logcatchsys/logtailf"
	"sync"

	"github.com/spf13/viper"
)

var mainOnce sync.Once
var configMgr map[string]*logconfig.ConfigData

const KEYCHANSIZE = 20

func ConstructMgr(configPaths interface{}, keyChan chan string, kafkaProducer *kafkaqueue.ProducerKaf) {

	for _, configData := range configPaths.([]interface{}) {
		conKey := ""
		conVal := ""
		for ckey, cval := range configData.(map[interface{}]interface{}) {
			if ckey == "logtopic" {
				conKey = cval.(string)
				continue
			}
			if ckey == "logpath" {
				conVal = cval.(string)
				continue
			}
		}
		if conKey == "" || conVal == "" {
			continue
		}
		configData := new(logconfig.ConfigData)
		configData.ConfigKey = conKey
		configData.ConfigValue = conVal
		ctx, cancel := context.WithCancel(context.Background())
		configData.ConfigCancel = cancel
		configMgr[configData.ConfigKey] = configData
		go logtailf.WatchLogFile(configData.ConfigKey, configData.ConfigValue,
			ctx, keyChan, kafkaProducer)
	}

}

func main() {
	v := viper.New()
	configPaths, confres := logconfig.ReadConfig(v)
	if configPaths == nil || !confres {
		fmt.Println("read config failed")
		return
	}

	producer, err := kafkaqueue.CreateKafkaProducer()
	if err != nil {
		fmt.Println("create producer failed ")
		return
	}

	kafkaProducer := &kafkaqueue.ProducerKaf{Producer: producer}

	configMgr = make(map[string]*logconfig.ConfigData)
	keyChan := make(chan string, KEYCHANSIZE)
	ConstructMgr(configPaths, keyChan, kafkaProducer)
	ctx, cancel := context.WithCancel(context.Background())
	pathChan := make(chan interface{})
	go logconfig.WatchConfig(ctx, v, pathChan)
	defer func() {
		mainOnce.Do(func() {
			if err := recover(); err != nil {
				fmt.Println("main goroutine panic ", err) // 这里的err其实就是panic传入的内容
			}
			cancel()
			for _, oldval := range configMgr {
				oldval.ConfigCancel()
			}
			configMgr = nil
			kafkaProducer.Producer.Close()
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
			pathDataNew := make(map[string]string)
			for _, configData := range pathData.([]interface{}) {
				conKey := ""
				conVal := ""
				for ckey, cval := range configData.(map[interface{}]interface{}) {
					if ckey == "logtopic" {
						conKey = cval.(string)
						continue
					}
					if ckey == "logpath" {
						conVal = cval.(string)
						continue
					}
				}
				if conKey == "" || conVal == "" {
					continue
				}
				pathDataNew[conKey] = conVal
			}

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
					configData.ConfigValue = conval
					ctx, cancel := context.WithCancel(context.Background())
					configData.ConfigCancel = cancel
					configMgr[conkey] = configData
					fmt.Println(conval)
					go logtailf.WatchLogFile(configData.ConfigKey, configData.ConfigValue,
						ctx, keyChan, kafkaProducer)
					continue
				}

				if oldval.ConfigValue != conval {
					oldval.ConfigValue = conval
					oldval.ConfigCancel()
					ctx, cancel := context.WithCancel(context.Background())
					oldval.ConfigCancel = cancel
					go logtailf.WatchLogFile(conkey, conval,
						ctx, keyChan, kafkaProducer)
					continue
				}

			}
			/*
				for mgrkey, mgrval := range configMgr {
					fmt.Println(mgrkey)
					fmt.Println(mgrval)
				}*/
		case keystr := <-keyChan:
			val, ok := configMgr[keystr]
			if !ok {
				continue
			}
			fmt.Println("recover goroutine watch ", keystr)
			var ctxcover context.Context
			ctxcover, val.ConfigCancel = context.WithCancel(context.Background())
			go logtailf.WatchLogFile(keystr, val.ConfigValue,
				ctxcover, keyChan, kafkaProducer)
		}
	}
}
