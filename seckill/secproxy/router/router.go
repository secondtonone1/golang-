package router

import (
	"golang-/seckill/secproxy/controller"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/seckill", &controller.SecKillCtr{}, "*:SecKill")
	beego.Router("/secinfo", &controller.SecKillCtr{}, "*:SecInfo")

}
