package etcdconsumer

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	etcdlogconf "golang-/logcatchsys/etcdlogconf"
	"golang-/logcatchsys/logconfig"
	"reflect"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func InitEtcdClient() (*clientv3.Client, error) {
	v := logconfig.InitVipper()
	if v == nil {
		fmt.Println("vipper init failed!")
		return nil, errors.New("vipper init failed!")
	}

	etcdconfig, etcdconfres := logconfig.ReadConfig(v, "etcdconfig")
	if !etcdconfres {
		fmt.Println("read config etcdconfig failed")
		return nil, errors.New("read config etcdconfig failed")
	}
	//fmt.Println(etcdKeys, " ", etcdconfig)
	endPoints := make([]string, 0, 20)
	for _, addr := range etcdconfig.([]interface{}) {
		endPoints = append(endPoints, addr.(string))
	}

	if len(endPoints) == 0 {
		return nil, errors.New("etcd addr is empty")
	}

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endPoints,
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		fmt.Println("connect failed, err:", err)
		return nil, errors.New("connect etcd failed!!")
	}

	return cli, nil
}

func GetTopicSet(cli *clientv3.Client) (interface{}, error) {
	etcdKeys, etcdres := logconfig.ReadConfig(logconfig.InitVipper(), "etcdkeys")
	if !etcdres {
		fmt.Println("read config etcdkeys failed")
		return nil, errors.New("read config etcdkeys failed")
	}
	fmt.Println(reflect.TypeOf(etcdKeys))
	topicSet := make(map[string]bool)
	for _, keyval := range etcdKeys.([]interface{}) {
		ctxtime, cancel := context.WithTimeout(context.Background(), time.Second)
		resp, err := cli.Get(ctxtime, keyval.(string))
		cancel()
		if err != nil {
			fmt.Println("get failed, err:", err)
			continue
		}

		for _, ev := range resp.Kvs {
			fmt.Printf("%s : %s ...\n", ev.Key, ev.Value)
			etcdLogConf := make([]*etcdlogconf.EtcdLogConf, 0, 20)
			unmarsherr := json.Unmarshal(ev.Value, &etcdLogConf)
			if unmarsherr != nil {
				fmt.Println("unmarshal error !, error is ", unmarsherr)
				continue
			}

			for _, etcdval := range etcdLogConf {
				topicSet[etcdval.Topic] = true
			}
		}

	}

	return topicSet, nil

}
