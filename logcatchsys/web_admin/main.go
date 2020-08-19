package main

import (
	"golang-/logcatchsys/web_admin/components"
	_ "golang-/logcatchsys/web_admin/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
)

func init_components() bool {
	err := components.InitLogger() //调用logger初始化
	if err != nil {
		logs.Warn("initDb failed, err :%v", err)
		return false
	}

	err = components.InitDb()
	if err != nil {
		logs.Warn("initDb failed, err:%v", err)
		return false
	}

	err = components.InitEtcd()
	if err != nil {
		logs.Warn("init etcd failed, err:%v", err)
		return false
	}

	return true
}

func main() {
	if init_components() == false {
		return
	}
	beego.Run()
}
