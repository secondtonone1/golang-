package logconfig

import (
	"context"
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func ReadConfig(v *viper.Viper) (interface{}, bool) {
	//设置读取的配置文件
	v.SetConfigName("config")
	//添加读取的配置文件路径
	v.AddConfigPath("./")
	//设置配置文件类型
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("err:%s\n", err)
		return nil, false
	}

	configPaths := v.Get("configpath")
	if configPaths == nil {
		return nil, false
	}

	return configPaths, true
}

func WatchConfig(ctx context.Context, v *viper.Viper, pathChan chan interface{}) {

	defer func() {
		fmt.Println("watch config goroutine exit")
		if err := recover(); err != nil {
			fmt.Println("watch config goroutine panic ", err)
		}
	}()

	//设置监听回调函数
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("config is change :%s \n", e.String())
		configPaths := v.Get("configpath")
		if configPaths == nil {
			return
		}
		pathChan <- configPaths
	})
	//开始监听
	v.WatchConfig()
	//信道不会主动关闭，可以主动调用cancel关闭
	<-ctx.Done()
}
