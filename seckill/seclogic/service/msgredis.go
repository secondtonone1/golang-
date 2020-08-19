package service

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"golang-/seckill/config"
	"golang-/seckill/seclogic/components"
	"sync"

	"github.com/astaxie/beego/logs"
	"github.com/garyburd/redigo/redis"
)

var MsgRdMgr *MsgRedisMgr
var gwait *sync.WaitGroup = new(sync.WaitGroup)

const (
	MSGCHANSIZE = 10000
)

func init() {
	MsgRdMgr = &MsgRedisMgr{
		MsgChanToRedis:   make(chan *MsgReqToRedis, MSGCHANSIZE),
		MsgChanFromRedis: make(chan *MsgRspFromRedis, MSGCHANSIZE),
		ToRedisGrClose:   make(chan struct{}),
		ToRedisWait:      new(sync.WaitGroup),
		FromRedisWait:    new(sync.WaitGroup),
		FromRedisGrClose: make(chan struct{}),
		NotifyMainClose:  make(chan struct{}),
	}

}

func InitRedisRWGoroutine() {
	gwait.Add(2)
	go WatchWriteGo(gwait)
	go WatchReadGo(gwait)
	gwait.Wait()
}

type MsgRedisMgr struct {
	MsgChanToRedis   chan *MsgReqToRedis
	MsgChanFromRedis chan *MsgRspFromRedis
	ToRedisGrClose   chan struct{}   //写redis协程组退出
	ToRedisWait      *sync.WaitGroup //wait 管理写redis携程组
	FromRedisWait    *sync.WaitGroup //wait 管理读redis协程组
	FromRedisGrClose chan struct{}   //读redis协程组退出
	NotifyMainClose  chan struct{}   //service主协程通知读写携程退出
}

type MsgReqToRedis struct {
	ProductId int
	UserId    int
}

type MsgRspFromRedis struct {
	ProductId int
	Status    int
	UserId    int
	Token     string
	Message   string
	Left      int
}

func WatchWriteGo(gwait *sync.WaitGroup) {
	MsgRdMgr.ToRedisWait.Add(components.SKConfData.RedisWriteGoCount)
	defer func() {
		close(MsgRdMgr.ToRedisGrClose)
		gwait.Done()
	}()
	for i := 0; i < components.SKConfData.RedisWriteGoCount; i++ {
		go WriteToRedis(MsgRdMgr.ToRedisWait)
	}
	MsgRdMgr.ToRedisWait.Wait()
}

func WatchReadGo(gwait *sync.WaitGroup) {
	MsgRdMgr.FromRedisWait.Add(components.SKConfData.RedisReadGoCount)
	defer func() {
		close(MsgRdMgr.FromRedisGrClose)
		gwait.Done()
	}()
	for i := 0; i < components.SKConfData.RedisReadGoCount; i++ {
		go ReadFromRedis(MsgRdMgr.FromRedisWait)
	}
	MsgRdMgr.FromRedisWait.Wait()
}

//logic 向redis中写， 写入msgfromredis chan中
func WriteToRedis(wg *sync.WaitGroup) {
	conn := components.MsgReqPool.Get()
	defer func() {
		wg.Done()
		conn.Close()
	}()
	for {
		select {
		case msgfromrd, ok := <-MsgRdMgr.MsgChanFromRedis:
			if !ok {
				logs.Debug("msg chan from redis closed")
				return
			}
			reply, err := json.Marshal(msgfromrd)
			if err != nil {
				logs.Debug("json marshal failed")
				continue
			}
			_, err = conn.Do("rpush", "msgfromredis", string(reply))

			if err != nil {
				logs.Debug("rpush to msgtoredis failed  ...%s", err.Error())
				continue
			} else {
				logs.Debug("logic write goroutine put data %v into redis ", string(reply))
			}
		}
	}
}

func ReadFromRedis(wg *sync.WaitGroup) {
	conn := components.MsgReqPool.Get()
	defer func() {
		wg.Done()
		conn.Close()
	}()
	for {
		reply, err := conn.Do("blpop", "msgtoredis", 0)
		if err != nil {
			logs.Debug("pop from msgtoredis failed ...%s", err.Error())
			continue
		}
		kvarray, err := redis.Strings(reply, err)
		if err != nil {
			logs.Debug("msgtoredis string convert failed, %v", err.Error())
			continue
		}
		logs.Debug("read from redis msgtoredis , popvalue  is %v", kvarray)
		msgtord := new(MsgReqToRedis)
		err = json.Unmarshal([]byte(kvarray[1]), msgtord)
		if err != nil {
			logs.Warn("json unmarshal failed , err is : %v", err.Error())
			continue
		}

		//加锁逻辑处理
		//更新product 信息
		components.SKConfData.SecInfoRWLock.Lock()
		defer components.SKConfData.SecInfoRWLock.Unlock()
		//这里忽略判断逻辑，直接默认成功
		msgrsp := new(MsgRspFromRedis)
		secinfo, ok := components.SKConfData.SecInfoData[msgtord.ProductId]
		if !ok {
			logs.Debug("product id %v not found", msgtord.ProductId)
			msgrsp.Status = config.PRODUCT_ID_INVALID
			msgrsp.Message = "product id not found "
			msgrsp.ProductId = msgtord.ProductId
			msgrsp.Status = config.PRODUCT_ID_INVALID
			msgrsp.UserId = msgtord.UserId
		} else if secinfo.Left > 0 {
			components.SKConfData.SecInfoData[msgtord.ProductId].Left -= 1
			msgrsp.Left = components.SKConfData.SecInfoData[msgtord.ProductId].Left
			tokendata := fmt.Sprintf("%d:%d", msgtord.ProductId, msgtord.UserId)
			tokenstr := fmt.Sprintf("%x", md5.Sum([]byte(tokendata)))
			msgrsp.Token = tokenstr
			msgrsp.Message = "seckill success"
			msgrsp.ProductId = msgtord.ProductId
			msgrsp.Status = config.STATUS_SEC_SUCCESS
			msgrsp.UserId = msgtord.UserId
		}
		//将回包放在msgfromrd中
		select {
		case MsgRdMgr.MsgChanFromRedis <- msgrsp:
			logs.Debug("logic put msgrsp into fromredis chan ")
			continue
		}

	}

}
