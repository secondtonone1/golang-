package service

import (
	"golang-/seckill/config"
	"golang-/seckill/secproxy/components"
	"strconv"
	"time"

	"github.com/astaxie/beego/logs"
)

//抢购接口
func SecKill(req *config.SecRequest) (data map[string]interface{}, err error) {

	data = make(map[string]interface{}, components.INIT_INFO_SIZE)

	//用户cookie校验先屏蔽
	/*
		if err = userCheck(req); err != nil {
			data["code"] = config.AUTH_SIGN_CHECK_FAILED
			data["message"] = "user auth sign check failed "
			return
		}
	*/

	isblack := isInblacklist(req.UserId, req.ClientAddr)
	if isblack {
		return
	}
	frequency := FrequencyMgrInst.CalFrequency(req.UserId, req.SecTimeStamp)
	if frequency > components.SKConfData.FrequencyLimit {
		data["status"] = config.FREQUENCY_LIMIT
		data["message"] = "user sec visit frequency limit "
		conn := components.BlacklistPool.Get()
		defer conn.Close()
		_, errs := conn.Do("rpush", "idblackqueue", strconv.Itoa(req.UserId))
		if errs != nil {
			logs.Debug("rpush idblackqueue err is %v", errs.Error())
		}
		return
	}

	ipfrequency := FrequencyMgrInst.CalIPFrequency(req.ClientAddr, req.SecTimeStamp)
	if ipfrequency > components.SKConfData.IpLimit {
		data["status"] = config.FREQUENCY_LIMIT
		data["message"] = "ip sec visit frequency limit "
		conn := components.BlacklistPool.Get()
		defer conn.Close()
		_, errs := conn.Do("rpush", "ipblackqueue", req.ClientAddr)
		if errs != nil {
			logs.Debug("rpush ipblackqueue err is %v", errs.Error())
		}
		return
	}

	msgtoredis := &MsgReqToRedis{
		ProductId: req.ProductId,
		UserId:    req.UserId,
	}
	writetimer := time.NewTimer(time.Second * 5)
	defer writetimer.Stop()
	select {
	case <-writetimer.C:
		logs.Debug("msg write to redis chan timeout, maybe chan has beeen closed")
		data["status"] = config.MSG_CHAN_CLOSED
		data["message"] = "msg chan to redis closed"
		return
	case MsgRdMgr.MsgChanToRedis <- msgtoredis:
		logs.Debug("msg chan to redis success")
	}

	//设置定时器，超时检测，防止请求阻塞
	ticker := time.NewTicker(time.Duration(10) * time.Second)
	defer func() {
		ticker.Stop()
	}()
	select {
	case msgrsp, ok := <-MsgRdMgr.MsgChanFromRedis:
		if !ok {
			logs.Debug("msg rsp from redis chan closed ")
			data["status"] = config.MSG_CHAN_CLOSED
			data["message"] = "msg chan from redis closed"
			return
		}

		if msgrsp.Status != config.STATUS_SEC_SUCCESS {
			logs.Debug(msgrsp.Message)
			data["status"] = msgrsp.Status
			data["message"] = msgrsp.Message
			return
		}
		datamap := make(map[string]interface{}, components.INIT_INFO_SIZE)
		datamap["productid"] = msgrsp.ProductId
		datamap["userid"] = msgrsp.UserId
		datamap["token"] = msgrsp.Token
		data["data"] = datamap
		data["status"] = config.STATUS_SEC_SUCCESS
		data["message"] = "seckill success"

		//更新product 信息
		components.SKConfData.SecInfoRWLock.Lock()
		defer components.SKConfData.SecInfoRWLock.Unlock()
		components.SKConfData.SecInfoData[msgrsp.ProductId].Left = msgrsp.Left
		return
	case <-ticker.C:
		data["status"] = config.STATUS_REQ_TIMEOUT
		data["message"] = "seckill timeout"
		return
	case <-MsgRdMgr.FromRedisGrClose:
		data["status"] = config.FROM_REDIS_GR_CLOSED
		data["message"] = "from redis chan group closed"
		return
	case <-MsgRdMgr.ToRedisGrClose:
		data["status"] = config.TO_REDIS_GR_CLOSED
		data["message"] = "to redis chan group closed"
		return
	}

}
