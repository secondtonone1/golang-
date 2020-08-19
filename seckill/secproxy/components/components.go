package components

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"golang-/seckill/config"
	"strconv"
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
		IDBlacklist: make(map[int]bool, INIT_INFO_SIZE),
		IPBlacklist: make(map[string]bool, INIT_INFO_SIZE),
	}
	BlacklistPool *redis.Pool
	EtcdClient    *etcdclient.Client
	MsgReqPool    *redis.Pool
)

func init() {
	err := loadConfig()
	if err != nil {
		logs.Debug("load config failed")
		panic("load config failed")
	}

	err = initRedisBlacklist()
	if err != nil {
		logs.Debug("initRedis failed")
		panic("initRedis failed")
	}

	err = initMsgReqPool()
	if err != nil {
		logs.Debug("init Msg Req Redis Pool failed")
		panic("init Msg Req Redis Pool failed")
	}

	err = loadBlacklist()
	if err != nil {
		logs.Debug("load black list failed")
		panic("laod redis black list failed ")
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
	logs.Debug("begin read redis_blacklist_addr config")
	RedisAddr := beego.AppConfig.String("redis_blacklist_addr")
	EtcdAddr := beego.AppConfig.String("etcd_addr")
	if len(RedisAddr) == 0 || len(EtcdAddr) == 0 {
		err = fmt.Errorf("read redisaddr[%s] or etcdaddr[%s] failed", RedisAddr, EtcdAddr)
		return
	}
	logs.Debug("begin read etcd_sec_prefix config")
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
	logs.Debug("begin read redis_blacklist_max_idle config")
	RMaxIdle, err := beego.AppConfig.Int("redis_blacklist_max_idle")
	if err != nil {
		err = fmt.Errorf("read redis_max_idle failed, error is [%v]", err)
		return
	}

	RMaxActive, err := beego.AppConfig.Int("redis_blacklist_max_active")
	if err != nil {
		err = fmt.Errorf("read redis_max_active failed, error is [%v]", err)
		return
	}

	RIdleTimeout, err := beego.AppConfig.Int("redis_blacklist_idle_timeout")
	if err != nil {
		err = fmt.Errorf("read redis_idle_timeout failed, error is [%v]", err)
		return
	}
	logs.Debug("begin read etcd time out config")
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

	CookieSecretKey := beego.AppConfig.String("cookie_secretkey")
	if len(CookieSecretKey) == 0 {
		err = fmt.Errorf("cookie secret key read failed")
		return
	}

	FrequencyLimit, err := beego.AppConfig.Int("frequency_limit")
	if err != nil {
		err = fmt.Errorf("frequency limit read failed")
		return
	}
	//logs.Debug("RedisAddr is %s", RedisAddr)
	//logs.Debug("EtcdAddr is %s", EtcdAddr)

	ReferList := beego.AppConfig.String("refer_whitelist")
	if len(ReferList) == 0 {
		err = fmt.Errorf("refer list read failed ")
		return
	}

	IpLimit, err := beego.AppConfig.Int("ip_limit")
	if err != nil {
		err = fmt.Errorf("read ip limit config failed ")
		return
	}

	SKConfData.RedisBlacklist.RedisAddr = RedisAddr
	SKConfData.RedisBlacklist.RedisIdleTime = RIdleTimeout
	SKConfData.RedisBlacklist.RedisMaxActive = RMaxActive
	SKConfData.RedisBlacklist.RedisMaxIdle = RMaxIdle
	SKConfData.EtcdConfData.EtcdAddr = EtcdAddr
	SKConfData.EtcdConfData.EtcdTimeout = EtcdTimeout
	SKConfData.EtcdConfData.EtcdSecPrefix = EtcdSecPrefix
	SKConfData.EtcdConfData.EtcdSecProduct = EtcdSecProduct
	SKConfData.LogConfData.LogLv = convertLogLv(LogLv)
	SKConfData.LogConfData.LogPath = LogPath
	SKConfData.LogConfData.MaxLines = LogLines
	SKConfData.CookieSecretKey = CookieSecretKey
	SKConfData.FrequencyLimit = FrequencyLimit
	SKConfData.ReferWhitelist = strings.Split(ReferList, ",")
	SKConfData.IpLimit = IpLimit

	read_redis_gnum, err := beego.AppConfig.Int("read_redis_count")
	if err != nil {
		err = fmt.Errorf("read read_redis_count failed ")
		return
	}

	SKConfData.RedisReadGoCount = read_redis_gnum
	write_redis_gnum, err := beego.AppConfig.Int("write_redis_count")
	if err != nil {
		err = fmt.Errorf("read write_redis_count failed ")
		return
	}

	SKConfData.RedisWriteGoCount = write_redis_gnum
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

func initRedisBlacklist() (err error) {
	BlacklistPool = &redis.Pool{
		MaxIdle:     SKConfData.RedisBlacklist.RedisMaxIdle,
		MaxActive:   SKConfData.RedisBlacklist.RedisMaxActive,
		IdleTimeout: time.Second * time.Duration(SKConfData.RedisBlacklist.RedisIdleTime),
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", SKConfData.RedisBlacklist.RedisAddr)
		},
	}

	conn := BlacklistPool.Get()
	defer conn.Close()
	_, err = conn.Do("ping")
	if err != nil {
		logs.Debug("redis ping failed, err is %v", err)
		return
	}

	return nil
}

func initMsgReqPool() (err error) {
	MsgReqPool = &redis.Pool{
		MaxIdle:     SKConfData.RedisBlacklist.RedisMaxIdle,
		MaxActive:   SKConfData.RedisBlacklist.RedisMaxActive,
		IdleTimeout: time.Second * time.Duration(SKConfData.RedisBlacklist.RedisIdleTime),
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", SKConfData.RedisBlacklist.RedisAddr)
		},
	}
	return nil
}

func loadBlacklist() error {
	conn := BlacklistPool.Get()
	SKConfData.BlacklistRWLock.Lock()
	defer func() {
		conn.Close()
		SKConfData.BlacklistRWLock.Unlock()
	}()

	/*
		replyres, err := conn.Do("hgetall", "idblacklist")
		if err != nil {
			logs.Warn("redis get id black list failed ")
			return errors.New("redis get id black list failed ")
		}
	*/

	replyres, err := conn.Do("hkeys", "idblacklist")
	if err != nil {
		logs.Warn("redis get id black list failed ")
		return errors.New("redis get id black list failed ")
	}

	idblacklist, err := redis.Strings(replyres, err)
	if err != nil {
		logs.Warn("redis convert command res to string slice error")
		return errors.New("redis convert command res to string slice error")
	}
	logs.Debug("load idblack list is %v", idblacklist)
	for _, blackid := range idblacklist {
		nid, err := strconv.Atoi(blackid)
		if err != nil {
			logs.Warn("id str convert int failed ,id is %v", nid)
			continue
		}
		SKConfData.IDBlacklist[nid] = true
	}
	/*
		iplist, err := conn.Do("hgetall", "ipblacklist")
		if err != nil {
			logs.Warn("redis get ip black list failed ")
			return errors.New("redis get ip black list failed ")
		}
	*/

	iplist, err := conn.Do("hkeys", "ipblacklist")
	if err != nil {
		logs.Warn("redis get ip black list failed ")
		return errors.New("redis get ip black list failed ")
	}

	ipblacklist, err := redis.Strings(iplist, err)
	if err != nil {
		logs.Warn("redis convert command res to string slice error")
		return errors.New("redis convert command res to string slice error")
	}
	logs.Debug("load ipblack list is %v", ipblacklist)
	for _, blackip := range ipblacklist {
		SKConfData.IPBlacklist[blackip] = true
	}

	go SyncBlackIdlist()
	go SyncBlackIplist()
	return nil
}

func SyncBlackIplist() {
	conn := BlacklistPool.Get()
	defer conn.Close()
	ipblacks := make(map[string]bool, 10)
	for {
		reply, err := conn.Do("blpop", "ipblackqueue", 0)
		if err != nil {
			logs.Debug("pop from ipblackqueue failed ...%s", err.Error())
			continue
		}
		kvarray, err := redis.Strings(reply, err)
		if err != nil {
			logs.Debug("ip black redis string convert failed, %v", err.Error())
			continue
		}
		logs.Debug("read from redis ipblackqueue , ip is %v", kvarray)
		ipblacks[kvarray[1]] = true
		/*
			if len(ipblacks) < 10 {
				continue
			}
		*/
		SKConfData.BlacklistRWLock.Lock()
		for ipkey, _ := range ipblacks {
			SKConfData.IPBlacklist[ipkey] = true
		}
		SKConfData.BlacklistRWLock.Unlock()
		for ipkey, _ := range ipblacks {
			conn.Do("hset", "ipblacklist", ipkey, ipkey)
		}
		ipblacks = make(map[string]bool, 10)
	}
}

func SyncBlackIdlist() {
	conn := BlacklistPool.Get()
	defer conn.Close()
	idblacks := make(map[string]bool, 10)
	for {
		reply, err := conn.Do("blpop", "idblackqueue", 0)
		if err != nil {
			logs.Debug("pop from idblackqueue failed  ...%s", err.Error())
			continue
		}
		kvarray, err := redis.Strings(reply, err)
		if err != nil {
			logs.Debug("id black redis string convert failed, %v", err.Error())
			continue
		}
		logs.Debug("read from redis idblackqueue , id is %v", kvarray)
		idblacks[kvarray[1]] = true
		/*
			if len(idblacks) < 10 {
				continue
			}
		*/
		SKConfData.BlacklistRWLock.Lock()
		for idkey, _ := range idblacks {
			nid, err := strconv.Atoi(idkey)
			if err != nil {
				logs.Debug("type convert failed ")
				continue
			}
			SKConfData.IDBlacklist[nid] = true
		}
		SKConfData.BlacklistRWLock.Unlock()
		for idkey, _ := range idblacks {
			conn.Do("hset", "idblacklist", idkey, idkey)
		}
		idblacks = make(map[string]bool, 10)
	}
}

func ReleaseRsc() {
	if BlacklistPool != nil {
		err := BlacklistPool.Close()
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
	SKConfData.SecInfoRWLock.Lock()
	defer SKConfData.SecInfoRWLock.Unlock()
	for _, secinfo := range secinfolist {
		if _, ok := SKConfData.SecInfoData[secinfo.ProductId]; ok {
			continue
		}
		sectmp := secinfo
		SKConfData.SecInfoData[secinfo.ProductId] = &sectmp
	}

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
		SKConfData.SecInfoRWLock.Lock()
		defer SKConfData.SecInfoRWLock.Unlock()
		for _, secinfo := range secinfolist {
			logs.Debug("secinfo.EndTime: %v\n", secinfo.EndTime)
			logs.Debug("secinfo.ProductId: %v\n", secinfo.ProductId)
			logs.Debug("secinfo.Total: %v\n", secinfo.Total)
			logs.Debug("secinfo.Status: %v\n", secinfo.Status)
			secinfocp := secinfo
			secinfocp.Left = secinfo.Total
			SKConfData.SecInfoData[secinfo.ProductId] = &secinfocp
		}
	}
	return
}
