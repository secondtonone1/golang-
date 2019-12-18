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
var etcdData

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

func ConstructEtcd(etcdDatas interface{}, keyChan chan string, kafkaProducer *kafkaqueue.ProducerKaf) {
	logkeys := etcdDatas.([]interface{})
	for _, logkey := range logkeys {
		fmt.Println(logkey.(string))
	}

}

//根据yaml文件修改后返回的配置信息，启动和关闭goroutine
func updateConfigGoroutine(pathData interface{}) {

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
	//删除监控日志
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
		//新增监控日志
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
		//修改监控日志
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
}

//根据etcd中的日志监控信息启动和关闭协程
func updateEtcdGoroutine(etcdlogData interface{}) {
	logkeys := etcdlogData.([]interface{})
	for _, logkey := range logkeys {
		fmt.Println(logkey.(string))
	}
}

func main() {
	v := viper.New()
	configPaths, confres := logconfig.ReadConfig(v, "collectlogs")
	if configPaths == nil || !confres {
		fmt.Println("read config failed")
		return
	}

	producer, err := kafkaqueue.CreateKafkaProducer()
	if err != nil {
		fmt.Println("create producer failed ")
		return
	}
	//构造协程监控配置中的日志
	kafkaProducer := &kafkaqueue.ProducerKaf{Producer: producer}
	configMgr = make(map[string]*logconfig.ConfigData)
	keyChan := make(chan string, KEYCHANSIZE)
	ConstructMgr(configPaths, keyChan, kafkaProducer)
	//监听配置文件
	ctx, cancel := context.WithCancel(context.Background())
	pathChan := make(chan interface{})
	etcdChan := make(chan interface{})
	go logconfig.WatchConfig(ctx, v, pathChan, etcdChan)
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
			updateConfigGoroutine(pathData)

		case etcdLogData, ok := <-etcdChan:
			if !ok {
				return
			}
			updateEtcdGoroutine(etcdLogData)
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
