package service

import (
	"encoding/json"
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

	initRedisRWGoroutine()
}

func initRedisRWGoroutine() {
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

//proxy向redis中写
func WriteToRedis(wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()
	for {
		select {
		case msgtoredis, ok := <-MsgRdMgr.MsgChanToRedis:
			if !ok {
				logs.Debug("msg chan to redis closed")
				return
			}
			_, err := json.Marshal(msgtoredis)
			if err != nil {
				logs.Debug("json marshal failed")
				continue
			}
			conn := components.MsgReqPool.Get()
			defer conn.Close()
			_, err = conn.Do("rpush", "msgtoredis", 0)

			if err != nil {
				logs.Debug("rpush to msgtoredis failed  ...%s", err.Error())
				continue
			}
		case <-MsgRdMgr.NotifyMainClose:
			logs.Debug("service main goroutine notify closed")
			return
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
		reply, err := conn.Do("blpop", "msgfromredis", 0)
		if err != nil {
			logs.Debug("pop from msgfromredis failed ...%s", err.Error())
			continue
		}
		kvarray, err := redis.Strings(reply, err)
		if err != nil {
			logs.Debug("msgfromredis string convert failed, %v", err.Error())
			continue
		}
		logs.Debug("read from redis msgfromredis , ip is %v", kvarray)
		msgfromrd := new(MsgRspFromRedis)
		err = json.Unmarshal([]byte(kvarray[1]), msgfromrd)
		if err != nil {
			logs.Warn("json unmarshal failed , err is : %v", err.Error())
			continue
		}
		select {
		case MsgRdMgr.MsgChanFromRedis <- msgfromrd:
			continue
		case <-MsgRdMgr.NotifyMainClose:
			logs.Debug("receive service main routine notify close")
			return
		}

	}

}
