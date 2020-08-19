package main

import (
	"context"
	"encoding/json"
	"fmt"
	"golang-/seckill/config"
	"strings"
	"time"

	bgconfig "github.com/astaxie/beego/config"
	"github.com/astaxie/beego/logs"
	"go.etcd.io/etcd/clientv3"
)

var (
	skconfdata config.SecKillConf
)

func set_product_etcd() (err error) {
	logs.Debug("set_product_etcd")
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{skconfdata.EtcdConfData.EtcdAddr},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		fmt.Println("connect failed, err:", err)
		logs.Debug("init etcd is ", err)
		return
	}
	logs.Debug("etcdkey prefix is %v", skconfdata.EtcdConfData.EtcdSecPrefix)
	lastsep := strings.HasSuffix(skconfdata.EtcdConfData.EtcdSecPrefix, "/")
	if !lastsep {
		skconfdata.EtcdConfData.EtcdSecPrefix = skconfdata.EtcdConfData.EtcdSecPrefix + "/"
	}

	secProductKey := skconfdata.EtcdConfData.EtcdSecPrefix + "product"
	secinfolist := []config.SecInfoConf{
		config.SecInfoConf{
			ProductId: 1001,
			StartTime: 1587350782,
			EndTime:   1587353781,
			Status:    0,
			Total:     100,
			Left:      100,
		},
		config.SecInfoConf{
			ProductId: 1002,
			StartTime: 1587350782,
			EndTime:   1587353781,
			Status:    0,
			Total:     100,
			Left:      100,
		},
	}

	jsonpro, err := json.Marshal(secinfolist)
	if err != nil {
		logs.Debug("pro marshal failed, err is %v ", err)
		return
	}

	putres, err := cli.Put(context.Background(), secProductKey, string(jsonpro))
	if err != nil {
		logs.Debug("cli put failed, err is %v", err)
		return
	}

	logs.Debug("putres is %v ", putres)
	return
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
	bgconfig, err := bgconfig.NewConfig("ini", "app.conf")
	if err != nil {
		logs.Error("init log failed, err is %v", err)
		return
	}

	logline, err := bgconfig.Int("log_maxlines")
	if err != nil {
		logs.Error("init log line failed %v ", err)
		return
	}

	loglv := bgconfig.String("log_level")
	if len(loglv) == 0 {
		logs.Error("init log level failed %v", loglv)
		return
	}

	logpath := bgconfig.String("log_path")
	if len(loglv) == 0 {
		logs.Error("init log level failed %v", logpath)
		return
	}
	logs.Debug("begin read etcd addr")
	EtcdAddr := bgconfig.String("etcd_addr")
	if len(EtcdAddr) == 0 {
		err = fmt.Errorf("read etcdaddr[%s] failed", EtcdAddr)
		return
	}

	logs.Debug("begin read etcd_sec_prefix ")
	EtcdSecPrefix := bgconfig.String("etcd_sec_prefix")
	if len(EtcdSecPrefix) == 0 {
		err = fmt.Errorf("read etcd_sec_prefix[%s] failed", EtcdSecPrefix)
		return
	}

	logs.Debug("begin read etcd_timeout ")
	EtcdTimeout, err := bgconfig.Int("etcd_timeout")
	if err != nil {
		err = fmt.Errorf("read etcd_time_out failed, error is [%v]", err)
		return
	}

	skconfdata.EtcdConfData.EtcdAddr = EtcdAddr
	skconfdata.EtcdConfData.EtcdSecPrefix = EtcdSecPrefix
	skconfdata.EtcdConfData.EtcdTimeout = EtcdTimeout
	skconfdata.LogConfData.LogLv = convertLogLv(loglv)
	skconfdata.LogConfData.LogPath = logpath
	skconfdata.LogConfData.MaxLines = logline

	logjson, err := json.Marshal(skconfdata.LogConfData)
	if err != nil {
		logs.Error("log json marshal failed , err is %v", err)
		return
	}

	logs.SetLogger(logs.AdapterFile, string(logjson))
	logs.SetLogFuncCall(true)
	return
}

func main() {
	err := loadConfig()
	if err != nil {
		panic(err.Error())
	}
	err = set_product_etcd()
	if err != nil {
		panic(err.Error())
	}
}
