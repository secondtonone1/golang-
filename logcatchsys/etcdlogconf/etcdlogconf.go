package etcdlogconf

import (
	"context"
	"encoding/json"
	"fmt"

	kafkaqueue "golang-/logcatchsys/kafka"
	logtailf "golang-/logcatchsys/logtailf"
	"time"

	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/mvcc/mvccpb"
)

const MaxChanSize = 20

var logConfChan chan string

func init() {

	logConfChan = make(chan string, MaxChanSize)
}

type EtcdLogConf struct {
	Path          string                  `json:"path"`
	Topic         string                  `json:"topic"`
	Ctx           context.Context         `json:"-"`
	Cancel        context.CancelFunc      `json:"-"`
	KeyChan       chan string             `json:"-"`
	KafkaProducer *kafkaqueue.ProducerKaf `json:"-"`
}

type EtcdLogMgr struct {
	Ctx           context.Context
	Cancel        context.CancelFunc
	KeyChan       chan string
	KafkaProducer *kafkaqueue.ProducerKaf
	EtcdKey       string
	EtcdClient    *clientv3.Client
	EtcdConfMap   map[string]*EtcdLogConf
}

func ConstructEtcd(etcdDatas interface{}, keyChan chan string, kafkaProducer *kafkaqueue.ProducerKaf) map[string]*EtcdLogMgr {
	etcdMgr := make(map[string]*EtcdLogMgr)
	if etcdDatas == nil {
		return etcdMgr
	}
	logkeys := etcdDatas.([]interface{})
	for _, logkey := range logkeys {
		clientv3 := InitEtcdClient()
		if clientv3 == nil {
			continue
		}
		etcdData := new(EtcdLogMgr)
		ctx, cancel := context.WithCancel(context.Background())
		etcdData.Ctx = ctx
		etcdData.Cancel = cancel
		etcdData.KafkaProducer = kafkaProducer
		etcdData.KeyChan = keyChan
		etcdData.EtcdKey = logkey.(string)
		etcdData.EtcdClient = clientv3
		etcdMgr[logkey.(string)] = etcdData
		fmt.Println(etcdData.EtcdKey, " init success ")
	}
	return etcdMgr
}

//根据etcd中的日志监控信息启动和关闭协程
func UpdateEtcdGoroutine(etcdMgr map[string]*EtcdLogMgr, etcdlogData interface{}, kafkaProducer *kafkaqueue.ProducerKaf, keyChan chan string) {
	if etcdlogData == nil {
		return
	}
	logkeys := etcdlogData.([]interface{})
	newkeyMap := make(map[string]bool)
	for _, logkey := range logkeys {
		fmt.Println("update key is ", logkey.(string))
		newkeyMap[logkey.(string)] = true
	}

	for oldkey, oldval := range etcdMgr {
		if _, ok := newkeyMap[oldkey]; !ok {
			oldval.Cancel()
			delete(etcdMgr, oldkey)
		}
	}

	for newkey, _ := range newkeyMap {
		if _, ok := etcdMgr[newkey]; !ok {
			clientv3 := InitEtcdClient()
			if clientv3 == nil {
				continue
			}
			etcdData := new(EtcdLogMgr)
			ctx, cancel := context.WithCancel(context.Background())
			etcdData.Ctx = ctx
			etcdData.Cancel = cancel
			etcdData.KafkaProducer = kafkaProducer
			etcdData.KeyChan = keyChan
			etcdData.EtcdKey = newkey
			etcdData.EtcdClient = clientv3
			etcdMgr[newkey] = etcdData
			fmt.Println(etcdData.EtcdKey, " init success ")
			go WatchEtcdKeys(etcdData)
		}
	}
}

func InitEtcdClient() *clientv3.Client {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:3379", "localhost:4379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return nil
	}
	//fmt.Println("connect succ")
	return cli
}

func WatchEtcdFile(etcdFile *EtcdLogConf) {
	logtailf.WatchLogFile(etcdFile.Topic, etcdFile.Path, etcdFile.Ctx, etcdFile.KeyChan, etcdFile.KafkaProducer)
}

func UpdateEtcdFile(etcdMgr *EtcdLogMgr, wresp *clientv3.WatchResponse) {
	etcdNewMap := make(map[string]*EtcdLogConf)
	for _, ev := range wresp.Events {
		fmt.Printf("%s %q:%q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		if ev.Type == mvccpb.DELETE {
			continue
		}
		//panic("test panic")
		etcdLogConfTmp := make([]*EtcdLogConf, 0, 20)
		unmarsherr := json.Unmarshal(ev.Kv.Value, &etcdLogConfTmp)
		if unmarsherr != nil {
			fmt.Println("unmarshal error !, error is ", unmarsherr)
			continue
		}
		for _, logslice := range etcdLogConfTmp {
			etcdNewMap[logslice.Topic] = logslice
		}
	}

	for oldkey, oldval := range etcdMgr.EtcdConfMap {
		_, ok := etcdNewMap[oldkey]
		if !ok {
			//该日志文件取消监控
			oldval.Cancel()
			delete(etcdMgr.EtcdConfMap, oldkey)
		}
	}

	for newkey, newval := range etcdNewMap {
		oldval, ok := etcdMgr.EtcdConfMap[newkey]
		if !ok {
			//新增日志文件，启动协程监控
			etcdMgr.EtcdConfMap[newval.Topic] = newval
			newval.Ctx, newval.Cancel = context.WithCancel(context.Background())
			newval.KeyChan = logConfChan
			newval.KafkaProducer = etcdMgr.KafkaProducer
			go WatchEtcdFile(newval)
			continue
		}

		//判断val是否修改
		if newval.Path != oldval.Path {
			oldval.Cancel()
			delete(etcdMgr.EtcdConfMap, oldval.Topic)
			etcdMgr.EtcdConfMap[newval.Topic] = newval
			newval.Ctx, newval.Cancel = context.WithCancel(context.Background())
			newval.KeyChan = logConfChan
			newval.KafkaProducer = etcdMgr.KafkaProducer
			go WatchEtcdFile(newval)
			continue
		}
	}
}

func WatchEtcdKeys(etcdMgr *EtcdLogMgr) {

	defer func() {
		if erreco := recover(); erreco != nil {
			etcdMgr.KeyChan <- etcdMgr.EtcdKey
			fmt.Println("watch etcd panic, exited")
			goto CLEARLOG_GOROUTINE
		}
		fmt.Println("watch etcd  exit")
		etcdMgr.EtcdClient.Close()
	CLEARLOG_GOROUTINE:
		for _, val := range etcdMgr.EtcdConfMap {
			val.Cancel()
		}
		etcdMgr.EtcdConfMap = nil
	}()
	etcdMgr.EtcdConfMap = make(map[string]*EtcdLogConf)
	ctxtime, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := etcdMgr.EtcdClient.Get(ctxtime, etcdMgr.EtcdKey)
	cancel()
	if err != nil {
		fmt.Println("get failed, err:", err)
		return
	}

	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s ...\n", ev.Key, ev.Value)
		etcdLogConf := make([]*EtcdLogConf, 0, 20)
		unmarsherr := json.Unmarshal(ev.Value, &etcdLogConf)
		if unmarsherr != nil {
			fmt.Println("unmarshal error !, error is ", unmarsherr)
			continue
		}

		for _, etcdval := range etcdLogConf {
			etcdMgr.EtcdConfMap[etcdval.Topic] = etcdval
			etcdval.Ctx, etcdval.Cancel = context.WithCancel(context.Background())
			etcdval.KeyChan = logConfChan
			etcdval.KafkaProducer = etcdMgr.KafkaProducer
			go WatchEtcdFile(etcdval)
		}
		fmt.Println(etcdMgr.EtcdConfMap)
	}
	watchChan := etcdMgr.EtcdClient.Watch(etcdMgr.Ctx, etcdMgr.EtcdKey)
	for {
		select {
		case wresp, ok := <-watchChan:
			if !ok {
				fmt.Println("watch etcd key  receive parent goroutine exited")
				return
			}
			UpdateEtcdFile(etcdMgr, &wresp)
		case logConfKey := <-logConfChan:
			etcdvalt, ok := etcdMgr.EtcdConfMap[logConfKey]
			if !ok {
				continue
			}
			//重启日志监控协程
			go WatchEtcdFile(etcdvalt)
		}
	}

}
