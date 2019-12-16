package logconfig

import (
	"context"
	"fmt"
	"path"
	"runtime"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var onceLogConf sync.Once

type ConfigData struct {
	ConfigKey    string
	ConfigValue  string
	ConfigCancel context.CancelFunc
}

func ReadConfig(v *viper.Viper) (interface{}, bool) {
	//设置读取的配置文件
	v.SetConfigName("config")
	//添加读取的配置文件路径
	_, filename, _, _ := runtime.Caller(0)
	//fmt.Println(filename)
	//fmt.Println(path.Dir(filename))
	v.AddConfigPath(path.Dir(filename))
	//设置配置文件类型
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("err:%s\n", err)
		return nil, false
	}

	collectlogs := v.Get("collectlogs")
	if collectlogs == nil {
		return nil, false
	}
	fmt.Println(collectlogs)
	return collectlogs, true
}

func WatchConfig(ctx context.Context, v *viper.Viper, pathChan chan interface{}) {

	defer func() {
		onceLogConf.Do(func() {
			fmt.Println("watch config goroutine exit")
			if err := recover(); err != nil {
				fmt.Println("watch config goroutine panic ", err)
			}
			close(pathChan)
		})
	}()

	//设置监听回调函数
	v.OnConfigChange(func(e fsnotify.Event) {
		//fmt.Printf("config is change :%s \n", e.String())
		collectlogs := v.Get("collectlogs")
		if collectlogs == nil {
			return
		}
		pathChan <- collectlogs
	})
	//开始监听
	v.WatchConfig()
	//信道不会主动关闭，可以主动调用cancel关闭
	<-ctx.Done()
}
