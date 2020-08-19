package components

import (
	"fmt"
	model "golang-/logcatchsys/web_admin/models"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func InitEtcd() (err error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{BConfig.String("etcd::etcd1"), BConfig.String("etcd::etcd2"),
			BConfig.String("etcd::etcd3")},
		DialTimeout: 5 * time.Second,
	})
	//logs.Debug("init etcd is ", err)
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}

	model.InitEtcd(cli)
	return
}
