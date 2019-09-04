package main

import (
	"fmt"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/broker/grpc"
	"github.com/micro/go-plugins/registry/etcdv3"
)

var (
	topic = "go.micro.web.topic.hi"
)

func main() {
	// 我这里用的etcd 做为服务发现，如果使用consul可以去掉
	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"localhost:2379", "localhost:3379", "localhost:4379",
		}
	})

	bk := grpc.NewBroker()

	err := bk.Init()
	if err != nil {
		fmt.Println("broker Init failed")
		return
	}

	err = bk.Connect()

	if err != nil {
		fmt.Println("broker Connect failed")
		return
	}
	defer bk.Disconnect()

	// 初始化服务
	service := micro.NewService(
		micro.Registry(reg),
		micro.Broker(bk),
	)
	service.Init()
	bk.Publish(topic, &broker.Message{

		Header: map[string]string{"type": "event"},

		Body: []byte(`an event`),
	})

}
