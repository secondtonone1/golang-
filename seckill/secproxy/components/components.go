package components

import (
	"context"
	"encoding/json"
	"fmt"
	"golang-/seckill/config"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/garyburd/redigo/redis"
	etcdclient "go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/mvcc/mvccpb"
)

const (
	INIT_INFO_SIZE = 100
)

var (
	SKConfData config.SecKillConf = config.SecKillConf{
		SecInfoData: make(map[int]*config.SecInfoConf, INIT_INFO_SIZE),
	}
	RedisPool  *redis.Pool
	EtcdClient *etcdclient.Client
)

func init() {
	err := loadConfig()
	if err != nil {
		logs.Debug("load config failed")
		panic("load config failed")
	}

	err = initRedis()
	if err != nil {
		logs.Debug("initRedis failed")
		panic("initRedis failed")
	}

	err = initEtcds()
	if err != nil {
		logs.Debug("initRedis failed")
		panic("initRedis failed")
	}

	err = initSecInfo()
	if err != nil {
		logs.Debug("initSecInfo failed")
		panic("initSecInfo failed")
	}
}

func convertLogLv(lvstr string) (lvint int) {
	switch lvstr {
	case "debug":
		return logs.LevelDebug
	case "warn":
		return logs.LevelWarn
	case "info":
		return logs.LevelInfo
	case "error":
		return logs.LevelError
	default:
		return logs.LevelDebug
	}
}

func loadConfig() (err error) {
	RedisAddr := beego.AppConfig.String("redis_addr")
	EtcdAddr := beego.AppConfig.String("etcd_addr")
	if len(RedisAddr) == 0 || len(EtcdAddr) == 0 {
		err = fmt.Errorf("read redisaddr[%s] or etcdaddr[%s] failed", RedisAddr, EtcdAddr)
		return
	}

	EtcdSecPrefix := beego.AppConfig.String("etcd_sec_prefix")
	if len(EtcdSecPrefix) == 0 {
		err = fmt.Errorf("read etcd_sec_prefix[%s] failed", EtcdSecPrefix)
		return
	}

	EtcdSecProduct := beego.AppConfig.String("etcd_sec_product")
	if len(EtcdSecProduct) == 0 {
		err = fmt.Errorf("read etcd_sec_product[%s] failed", EtcdSecProduct)
		return
	}

	RMaxIdle, err := beego.AppConfig.Int("redis_max_idle")
	if err != nil {
		err = fmt.Errorf("read redis_max_idle failed, error is [%v]", err)
		return
	}

	RMaxActive, err := beego.AppConfig.Int("redis_max_active")
	if err != nil {
		err = fmt.Errorf("read redis_max_active failed, error is [%v]", err)
		return
	}

	RIdleTimeout, err := beego.AppConfig.Int("redis_idle_timeout")
	if err != nil {
		err = fmt.Errorf("read redis_idle_timeout failed, error is [%v]", err)
		return
	}

	EtcdTimeout, err := beego.AppConfig.Int("etcd_timeout")
	if err != nil {
		err = fmt.Errorf("read etcd_time_out failed, error is [%v]", err)
		return
	}
	logs.Debug("begin read log config")
	LogLv := beego.AppConfig.String("log_level")
	if len(LogLv) == 0 {
		err = fmt.Errorf("read log_level[%s] failed", LogLv)
		return
	}

	LogPath := beego.AppConfig.String("log_path")
	if len(LogPath) == 0 {
		err = fmt.Errorf("read log_path[%s] failed", LogPath)
		return
	}
	logs.Debug("begin read log maxlines")
	LogLines, err := beego.AppConfig.Int("log_maxlines")
	if err != nil {
		err = fmt.Errorf("read logmaxlines failed, error is [%v]", err)
		return
	}

	//logs.Debug("RedisAddr is %s", RedisAddr)
	//logs.Debug("EtcdAddr is %s", EtcdAddr)

	SKConfData.RdisConfData.RedisAddr = RedisAddr
	SKConfData.RdisConfData.RedisIdleTime = RIdleTimeout
	SKConfData.RdisConfData.RedisMaxActive = RMaxActive
	SKConfData.RdisConfData.RedisMaxIdle = RMaxIdle
	SKConfData.EtcdConfData.EtcdAddr = EtcdAddr
	SKConfData.EtcdConfData.EtcdTimeout = EtcdTimeout
	SKConfData.EtcdConfData.EtcdSecPrefix = EtcdSecPrefix
	SKConfData.EtcdConfData.EtcdSecProduct = EtcdSecProduct
	SKConfData.LogConfData.LogLv = convertLogLv(LogLv)
	SKConfData.LogConfData.LogPath = LogPath
	SKConfData.LogConfData.MaxLines = LogLines
	logs.Debug("begin marshal log")
	logjson, err := json.Marshal(SKConfData.LogConfData)
	if err != nil {
		err = fmt.Errorf("json marshal failed, logjson err is %v", err)
		return
	}

	logs.SetLogger(logs.AdapterFile, string(logjson))
	logs.SetLogFuncCall(true)
	return nil
}

func initRedis() (err error) {
	RedisPool = &redis.Pool{
		MaxIdle:     SKConfData.RdisConfData.RedisMaxIdle,
		MaxActive:   SKConfData.RdisConfData.RedisMaxActive,
		IdleTimeout: time.Second * time.Duration(SKConfData.RdisConfData.RedisIdleTime),
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", SKConfData.RdisConfData.RedisAddr)
		},
	}

	conn := RedisPool.Get()
	defer conn.Close()
	_, err = conn.Do("ping")
	if err != nil {
		logs.Debug("redis ping failed, err is %v", err)
		return
	}

	return nil
}

func ReleaseRsc() {
	if RedisPool != nil {
		err := RedisPool.Close()
		if err != nil {
			logs.Error("Redis Pool release failed, err is %v", err)
		}
	}

	if EtcdClient != nil {
		err := EtcdClient.Close()
		if err != nil {
			logs.Error("Etcd Client release failed, err is %v", err)
		}
	}
}

func watchEtcd() {
	if EtcdClient == nil {
		panic("etcd client is nil")
	}

	lastsep := strings.HasSuffix(SKConfData.EtcdConfData.EtcdSecPrefix, "/")
	if !lastsep {
		SKConfData.EtcdConfData.EtcdSecPrefix = SKConfData.EtcdConfData.EtcdSecPrefix + "/"
	}

	secProductKey := SKConfData.EtcdConfData.EtcdSecPrefix + SKConfData.EtcdConfData.EtcdSecProduct

	watchchan := EtcdClient.Watch(context.Background(), secProductKey)

	for wrsp := range watchchan {
		bupdate := false
		secinflist := make([]config.SecInfoConf, 0)
		for _, wevent := range wrsp.Events {
			if wevent.Type == mvccpb.DELETE {
				logs.Warn("key [%s] is deleted ", secProductKey)
				continue
			}

			if wevent.Type == mvccpb.PUT && string(wevent.Kv.Key) == secProductKey {
				jsonres := json.Unmarshal(wevent.Kv.Value, &secinflist)
				if jsonres != nil {
					logs.Error("json unmarshal failed, error is %v", jsonres)
					continue
				}
				bupdate = true
			}

		}

		if bupdate {
			updateSecInfoData(secinflist)
		}
	}
}

func updateSecInfoData(secinfolist []config.SecInfoConf) {
	secinfomap := make(map[int]*config.SecInfoConf, INIT_INFO_SIZE)
	for _, secinfo := range secinfolist {
		sectmp := secinfo
		secinfomap[secinfo.ProductId] = &sectmp
	}

	SKConfData.SecInfoRLock.Lock()
	defer SKConfData.SecInfoRLock.Unlock()
	SKConfData.SecInfoData = secinfomap
	for key, val := range SKConfData.SecInfoData {
		logs.Debug("key is %d", key)
		logs.Debug("secinfo.EndTime: %v\n", val.EndTime)
		logs.Debug("secinfo.ProductId: %v\n", val.ProductId)
		logs.Debug("secinfo.Total: %v\n", val.Total)
		logs.Debug("secinfo.Status: %v\n", val.Status)
	}
}

func initEtcds() (err error) {
	cli, err := etcdclient.New(etcdclient.Config{
		Endpoints:   []string{SKConfData.EtcdConfData.EtcdAddr},
		DialTimeout: time.Duration(SKConfData.EtcdConfData.EtcdTimeout) * time.Second,
	})

	if err != nil {
		logs.Debug("init etcd failed, error is %v", err)
		return
	}

	EtcdClient = cli

	go watchEtcd()
	return
}

func initSecInfo() (err error) {
	lastsep := strings.HasSuffix(SKConfData.EtcdConfData.EtcdSecPrefix, "/")
	if !lastsep {
		SKConfData.EtcdConfData.EtcdSecPrefix = SKConfData.EtcdConfData.EtcdSecPrefix + "/"
	}

	secProductKey := SKConfData.EtcdConfData.EtcdSecPrefix + SKConfData.EtcdConfData.EtcdSecProduct
	proRsp, err := EtcdClient.Get(context.Background(), secProductKey)
	if err != nil {
		logs.Debug("etcd get key[%s] failed, err is %v", secProductKey, err)
		return
	}

	for prok, prov := range proRsp.Kvs {
		logs.Debug("prok is %s, prov is %s\n", prok, prov)
		secinfolist := make([]config.SecInfoConf, 0)
		if err = json.Unmarshal(prov.Value, &secinfolist); err != nil {
			logs.Error("json unmarshal failed, error is %v ", err)
			return
		}
		SKConfData.SecInfoRLock.Lock()
		defer SKConfData.SecInfoRLock.Unlock()
		for _, secinfo := range secinfolist {
			logs.Debug("secinfo.EndTime: %v\n", secinfo.EndTime)
			logs.Debug("secinfo.ProductId: %v\n", secinfo.ProductId)
			logs.Debug("secinfo.Total: %v\n", secinfo.Total)
			logs.Debug("secinfo.Status: %v\n", secinfo.Status)
			secinfocp := secinfo
			SKConfData.SecInfoData[secinfo.ProductId] = &secinfocp
		}
	}
	return
}
