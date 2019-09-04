package main

import (
	"context"
	"fmt"
	model "golang-/microdemo3/proto/model"
	rpcapi "golang-/microdemo3/proto/rpcapi"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"

	"time"

	"github.com/micro/go-micro/util/log"
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

	// New Service
	service := micro.NewService(
		micro.Name("lps.srv.eg1"),
		micro.Version("latest"),
		micro.Registry(reg),
	)

	// Initialise service
	service.Init()

	b := service.Client().Options().Broker

	if err := b.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}
	if err := b.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}
	defer b.Disconnect()

	go func(topic string, bk broker.Broker) {
		fmt.Println("broker publish msg")
		bk.Publish(topic, &broker.Message{
			Header: map[string]string{"type": "event"},

			Body: []byte("an event"),
		})

	}(topic, b)

	go func() {
		time.Sleep(time.Second * time.Duration(1))
		sayClent := rpcapi.NewSayService("lps.srv.eg1", service.Client())

		rsp, err := sayClent.Hello(context.Background(), &model.SayParam{Msg: "hello server"})
		if err != nil {
			panic(err)
		}

		fmt.Println(rsp)
	}()

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
