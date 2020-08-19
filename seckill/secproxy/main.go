package main

import (
	_ "golang-/seckill/secproxy/router"

	components "golang-/seckill/secproxy/components"

	"github.com/astaxie/beego"
)

func main() {
	defer func() {
		components.ReleaseRsc()
	}()
	beego.Run()
}
