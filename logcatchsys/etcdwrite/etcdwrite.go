package main

import (
	"context"
	"encoding/json"
	"fmt"
	logconfig "golang-/logcatchsys/logconfig"
	"time"

	"go.etcd.io/etcd/clientv3"
	//"github.com/etcd-io/etcd/clientv3"
)

const (
	EtcdKey = "collectlogkey1"
)

type LogConf struct {
	Path  string `json:"path"`
	Topic string `json:"topic"`
}

type EtcdClient struct {
	etcdClient *clientv3.Client
}

func InitEtcdConf() (*EtcdClient, bool) {
	etcdconf, _ := logconfig.ReadConfig(logconfig.InitVipper(), "etcdconfig")
	if etcdconf == nil {
		return nil, false
	}

	etcdaddrs := make([]string, 0, 20)
	for _, val := range etcdconf.([]interface{}) {
		etcdaddrs = append(etcdaddrs, val.(string))
	}

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   etcdaddrs,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return nil, false
	}
	fmt.Println("connect succ")
	return &EtcdClient{etcdClient: cli}, true
}

func (ec *EtcdClient) ReleaseEtcd() {
	ec.etcdClient.Close()
}

func (ec *EtcdClient) Put(ctx context.Context, etcdKey string, data string) (putres *clientv3.PutResponse, err error) {
	putres, err = ec.etcdClient.Put(ctx, etcdKey, data)
	return
}

func (ec *EtcdClient) Get(ctx context.Context, etcdKey string) (getres *clientv3.GetResponse, err error) {
	getres, err = ec.etcdClient.Get(ctx, etcdKey)
	return
}

func main() {
	etcdClient, bres := InitEtcdConf()
	if bres == false || etcdClient == nil {
		return
	}
	defer etcdClient.ReleaseEtcd()

	var logConfArr []LogConf

	logConfArr = append(
		logConfArr,
		LogConf{
			Path:  "D:/golangwork/src/golang-/logcatchsys/logdir1/log.txt",
			Topic: "golang_log",
		},
	)

	logConfArr = append(
		logConfArr,
		LogConf{
			Path:  "D:/golangwork/src/golang-/logcatchsys/logdir2/log.txt",
			Topic: "etcd_log",
		},
	)

	data, err := json.Marshal(logConfArr)
	if err != nil {
		fmt.Println("json failed, ", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = etcdClient.Put(ctx, EtcdKey, string(data))
	cancel()
	if err != nil {
		fmt.Println("put failed, err:", err)
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := etcdClient.Get(ctx, EtcdKey)
	cancel()
	if err != nil {
		fmt.Println("get failed, err:", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}

}
