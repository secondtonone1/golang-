package main

import (
	"context"
	"fmt"
	"path"
	"runtime"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func ReadConfig(v *viper.Viper) {
	//设置读取的配置文件
	v.SetConfigName("config")
	//添加读取的配置文件路径
	v.AddConfigPath("./")
	//设置配置文件类型
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("err:%s\n", err)
	}

	logrelative := v.Get("configpath.logdir1")

	_, filename, _, _ := runtime.Caller(0)
	fmt.Println(filename)
	datapath := path.Join(path.Dir(filename), logrelative.(string))
	fmt.Println("datapath is ", datapath)

}

func WatchConfig(v *viper.Viper) {
	//创建一个信道等待关闭（模拟服务器环境）
	ctx, _ := context.WithCancel(context.Background())
	//cancel可以关闭信道
	//ctx, cancel := context.WithCancel(context.Background())
	//设置监听回调函数
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("config is change :%s \n", e.String())
		configPaths := v.Get("configpath")
		if v.Get("configpath") == nil {
			return
		}
		fmt.Println(configPaths)
		for idx, val := range configPaths.(map[string]interface{}) {
			fmt.Println(idx)
			fmt.Println(val)
		}

		//cancel()
	})
	//开始监听
	v.WatchConfig()
	//信道不会主动关闭，可以主动调用cancel关闭
	<-ctx.Done()
}

func main() {

	v := viper.New()
	ReadConfig(v)
	WatchConfig(v)
}
