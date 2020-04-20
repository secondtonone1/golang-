package controller

import (
	"golang-/seckill/config"
	"golang-/seckill/secproxy/service"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type SecKillCtr struct {
	beego.Controller
}

func (skc *SecKillCtr) SecKill() {
	logs.Debug("receive SecKill ")
	skc.Data["json"] = "SecKill"
	skc.ServeJSON(true)
}

func (skc *SecKillCtr) SecInfo() {
	logs.Debug("receive SecInfo ")
	inforsp := make(map[string]interface{})
	defer func() {
		skc.Data["json"] = inforsp
		skc.ServeJSON(true)
	}()

	productid, err := skc.GetInt("productid")
	if err != nil {
		logs.Debug("productid not found, we get product list")
		data, err := service.GetProductList()
		if err != nil {
			inforsp["code"] = config.STATUS_PRODUCT_LIST_ERR
			inforsp["message"] = "get product info list failed"
			logs.Debug("get product faild , error code is %d", inforsp["code"])
			return
		}

		inforsp["code"] = config.STATUS_SELL_NORMAL
		inforsp["message"] = "get product info list success"
		inforsp["data"] = data
		return
	}

	data, err := service.GetProductById(productid)

	if err != nil {
		inforsp["code"] = data["status"]
		inforsp["message"] = data["message"]
		logs.Debug("get product faild , error code is %d", inforsp["code"])
		return
	}

	inforsp["code"] = data["status"]
	inforsp["message"] = data["message"]
	inforsp["data"] = []map[string]interface{}{data}

}
