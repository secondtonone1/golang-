package service

import (
	"crypto/md5"
	"errors"
	"fmt"
	"golang-/seckill/config"
	"golang-/seckill/secproxy/components"
	"time"

	"github.com/astaxie/beego/logs"
)

func convertProduct(data map[string]interface{}, product *config.SecInfoConf) error {
	defer func() {
		data["productid"] = product.ProductId
		data["starttime"] = time.Unix(product.StartTime, 0).Format("2006-01-02 03:04:05 PM")
		data["endtime"] = time.Unix(product.EndTime, 0).Format("2006-01-02 03:04:05 PM")
		data["total"] = product.Total
		data["left"] = product.Left
	}()

	if product.Status == config.STATUS_SELL_OUT {
		data["status"] = config.STATUS_SELL_OUT
		data["message"] = "product sell out"
		return nil
	}

	if product.Status == config.STATUS_SELL_FORBID {
		data["status"] = config.STATUS_SELL_FORBID
		data["message"] = "product forbid sell "
		return nil
	}

	curtime := time.Now().Unix()
	if curtime < product.StartTime {
		data["status"] = config.STATUS_SELL_NOT_BEGIN
		data["message"] = "product sell not begin"
		return nil
	}

	if curtime > product.EndTime {
		data["status"] = config.STATUS_SELL_END
		data["message"] = "product sell end"
		return nil
	}

	data["status"] = config.STATUS_SEC_SUCCESS
	data["message"] = "get secinfo success"

	return nil
}

func GetProductById(productid int) (data map[string]interface{}, err error) {
	data = make(map[string]interface{}, components.INIT_INFO_SIZE)
	components.SKConfData.SecInfoRWLock.RLock()
	product, ok := components.SKConfData.SecInfoData[productid]

	if !ok {
		logs.Debug("can't found product id [%d] ", productid)
		err = fmt.Errorf("product[%d] not found", productid)
		data["message"] = "product not found"
		data["status"] = config.STATUS_PRODUCT_NOT_FOUND
		return data, err
	}

	if product.EndTime < product.StartTime {
		err = errors.New("endtime can't be before starttime")
		data["message"] = "endtime litter than starttime"
		data["status"] = config.STATUS_PRODUCT_TIME_ERR
		return data, err
	}
	components.SKConfData.SecInfoRWLock.RUnlock()

	err = convertProduct(data, product)
	if err != nil {
		logs.Debug("can't convert product id [%d] ", productid)
		err = fmt.Errorf("can't convert product id [%d]", productid)
		data["message"] = "convert product failed"
		data["status"] = config.CONVERT_PRODUCT_INFO_ERR
		return data, err
	}

	return data, nil
}

//golang rwlock rlock可以嵌套使用，其他锁不可以嵌套，lock中不能嵌套lock，也不能嵌套rlock
//为了避免隐患，这里内部不调用getproductbyid，虽然嵌套调用rlock不会出问题
func GetProductList() (data []map[string]interface{}, err error) {
	data = make([]map[string]interface{}, 0)
	components.SKConfData.SecInfoRWLock.RLock()
	defer components.SKConfData.SecInfoRWLock.RUnlock()
	for _, secval := range components.SKConfData.SecInfoData {
		infomap := make(map[string]interface{}, components.INIT_INFO_SIZE)
		err := convertProduct(infomap, secval)
		if err != nil {
			continue
		}
		data = append(data, infomap)
	}
	return
}

func userCheck(req *config.SecRequest) (err error) {
	//检测跳转
	/*
		found := false
		for _, refer := range components.SKConfData.ReferWhitelist {
			if refer == req.ReferAddr {
				found = true
			}
		}

		if found == false {
			return errors.New("refer addr is invalid ")
		}
	*/
	authData := fmt.Sprintf("%d:%s", req.UserId, components.SKConfData.CookieSecretKey)
	authSign := fmt.Sprintf("%x", md5.Sum([]byte(authData)))
	if authSign != req.UserAuthSign {
		return errors.New("user auth sign dosen't match ")
	}

	return nil
}

//抢购接口
func SecKill(req *config.SecRequest) (data map[string]interface{}, err error) {
	components.SKConfData.SecInfoRWLock.RLock()
	data = make(map[string]interface{}, components.INIT_INFO_SIZE)
	defer func() {
		components.SKConfData.SecInfoRWLock.RUnlock()
	}()

	//用户cookie校验先屏蔽
	/*
		if err = userCheck(req); err != nil {
			data["code"] = config.AUTH_SIGN_CHECK_FAILED
			data["message"] = "user auth sign check failed "
			return
		}
	*/

	frequency := FrequencyMgrInst.CalFrequency(req.UserId, req.SecTimeStamp)
	if frequency > components.SKConfData.FrequencyLimit {
		data["status"] = config.FREQUENCY_LIMIT
		data["message"] = "user sec visit frequency limit "
		return
	}

	ipfrequency := FrequencyMgrInst.CalIPFrequency(req.ClientAddr, req.SecTimeStamp)
	if ipfrequency > components.SKConfData.IpLimit {
		data["status"] = config.FREQUENCY_LIMIT
		data["message"] = "ip sec visit frequency limit "
		return
	}

	data["status"] = config.STATUS_SEC_SUCCESS
	data["message"] = "seckill success"
	data["data"] = nil
	return
}
