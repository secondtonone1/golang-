package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
	//"github.com/etcd-io/etcd/clientv3"
)

const (
	EtcdKey = "watchkey"
)

type LogConf struct {
	Path  string `json:"path"`
	Topic string `json:"topic"`
}

type EtcdClient struct {
	etcdClient *clientv3.Client
}

func InitEtcdConf() (*EtcdClient, bool) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:3379", "localhost:4379"},
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

func (ec *EtcdClient) Watch(ctx context.Context, etcdKey string) clientv3.WatchChan {
	return ec.etcdClient.Watch(ctx, etcdKey)
}

func main() {
	etcdClient, bres := InitEtcdConf()
	if bres == false || etcdClient == nil {
		return
	}
	defer etcdClient.ReleaseEtcd()

	for {
		watchChan := etcdClient.Watch(context.Background(), EtcdKey)
		for wresp := range watchChan {
			for _, ev := range wresp.Events {
				fmt.Printf("%s %q:%q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			}
		}
	}

}
