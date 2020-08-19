package controller

import (
	"golang-/seckill/config"
	"golang-/seckill/secproxy/service"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type SecKillCtr struct {
	beego.Controller
}

func (skc *SecKillCtr) SecKill() {
	logs.Debug("receive SecKill ")
	productId, err := skc.GetInt("product_id")
	inforsp := make(map[string]interface{})
	defer func() {
		skc.Data["json"] = inforsp
		skc.ServeJSON(true)
	}()

	if err != nil {
		inforsp["code"] = config.STATUS_PRODUCTID_INVALID
		inforsp["message"] = "invalid product id "
		return
	}

	secReq := &config.SecRequest{}
	source := skc.GetString("src")
	authcode := skc.GetString("authcode")
	secTime := skc.GetString("time")
	nance := skc.GetString("nance")
	userIdInt, err := skc.GetInt("user_id")
	if err != nil {
		inforsp["code"] = config.USER_ID_INVALID
		inforsp["message"] = "invalid user id  "
		return
	}

	/*
		usrId := skc.Ctx.GetCookie("userId")
		userIdInt, err := strconv.Atoi(usrId)
		if err != nil {
			inforsp["code"] = config.STRING_CONVERT_FAILED
			inforsp["message"] = "string convert failed  "
			return
		}
	*/
	userAuthSign := skc.Ctx.GetCookie("userAuthSign")
	secReq.ProductId = productId
	secReq.Source = source
	secReq.AuthCode = authcode
	secReq.SecTime = secTime
	secReq.Nance = nance
	secReq.UserId = userIdInt
	secReq.UserAuthSign = userAuthSign
	secReq.SecTimeStamp = time.Now().Unix()
	//获取对方地址
	if len(skc.Ctx.Request.RemoteAddr) > 0 {
		lastindex := strings.LastIndex(skc.Ctx.Request.RemoteAddr, ":")
		secReq.ClientAddr = skc.Ctx.Request.RemoteAddr[:lastindex]
	}
	logs.Debug("client addr is %v", secReq.ClientAddr)
	//从哪个地址跳转过来的
	secReq.ReferAddr = skc.Ctx.Request.Referer()
	logs.Debug("refer addr is %v", secReq.ReferAddr)
	logs.Debug("secreq is [%v]", secReq)
	data, err := service.SecKill(secReq)
	if err != nil {
		inforsp["code"] = data["status"]
		inforsp["message"] = data["message"]
		return
	}

	inforsp["code"] = data["status"]
	inforsp["message"] = data["message"]
	inforsp["data"] = data

}

func (skc *SecKillCtr) SecInfo() {
	logs.Debug("receive SecInfo ")
	inforsp := make(map[string]interface{})
	defer func() {
		skc.Data["json"] = inforsp
		skc.ServeJSON(true)
	}()

	productid, err := skc.GetInt("product_id")
	if err != nil {
		logs.Debug("productid not found, we get product list")
		data, err := service.GetProductList()
		if err != nil {
			inforsp["code"] = config.STATUS_PRODUCT_LIST_ERR
			inforsp["message"] = "get product info list failed"
			logs.Debug("get product faild , error code is %d", inforsp["code"])
			return
		}

		inforsp["code"] = config.STATUS_SEC_SUCCESS
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
