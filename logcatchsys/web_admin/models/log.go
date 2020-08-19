package model

import (
	"context"
	"encoding/json"

	"time"

	etcdlogconf "golang-/logcatchsys/etcdlogconf"

	"github.com/astaxie/beego/logs"
	etcdclient "go.etcd.io/etcd/clientv3"
)

var (
	etcdClient *etcdclient.Client
)

type LogInfo struct {
	AppId      int    `db:"app_id"`
	AppName    string `db:"app_name"`
	LogId      int    `db:"log_id"`
	Status     int    `db:"status"`
	CreateTime string `db:"create_time"`
	LogPath    string `db:"log_path"`
	Topic      string `db:"topic"`
}

func InitEtcd(client *etcdclient.Client) {
	etcdClient = client
}

func GetAllLogInfo() (loglist []LogInfo, err error) {
	err = Db.Select(&loglist,
		"select a.app_id, b.app_name, a.create_time, a.topic, a.log_id, a.status, a.log_path from tbl_log_info a, tbl_app_info b where a.app_id=b.app_id")
	if err != nil {
		logs.Warn("Get All App Info failed, err:%v", err)
		return
	}
	return
}

func CreateLog(info *LogInfo) (err error) {

	conn, err := Db.Begin()
	if err != nil {
		logs.Warn("CreateApp failed, Db.Begin error:%v", err)
		return
	}

	defer func() {
		if err != nil {
			conn.Rollback()
			return
		}

		conn.Commit()
	}()

	var appId []int
	err = Db.Select(&appId, "select app_id from tbl_app_info where app_name=?", info.AppName)
	if err != nil || len(appId) == 0 {
		logs.Warn("select app_id failed, Db.Exec error:%v", err)
		return
	}

	timeStr := time.Now().Format("2006-01-02 15:04:05")
	info.AppId = appId[0]
	r, err := conn.Exec("insert into tbl_log_info(app_id, log_path, topic, app_name, create_time)values(?, ?, ?,?,?)",
		info.AppId, info.LogPath, info.Topic, info.AppName, timeStr)

	if err != nil {
		logs.Warn("CreateApp failed, Db.Exec error:%v", err)
		return
	}

	_, err = r.LastInsertId()
	if err != nil {
		logs.Warn("CreateApp failed, Db.LastInsertId error:%v", err)
		return
	}

	return
}

func SetLogConfToEtcd(etcdKey string, info *LogInfo) (err error) {

	var logConfArr []etcdlogconf.EtcdLogConf

	logConfArr = append(
		logConfArr,
		etcdlogconf.EtcdLogConf{
			Path:  info.LogPath,
			Topic: info.Topic,
		},
	)

	data, err := json.Marshal(logConfArr)
	if err != nil {
		logs.Warn("marshal failed, err:%v", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//cli.Delete(ctx, EtcdKey)
	//return
	_, err = etcdClient.Put(ctx, etcdKey, string(data))
	cancel()
	if err != nil {
		logs.Warn("Put failed, err:%v", err)
		return
	}

	logs.Debug("put etcd succ, data:%v", string(data))
	return
}
