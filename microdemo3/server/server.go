package main

import (
	"context"
	"fmt"
	model "golang-/microdemo3/proto/model"
	rpcapi "golang-/microdemo3/proto/rpcapi"

	"github.com/micro/go-micro/util/log"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
)

var (
	topic = "go.micro.web.topic.hi"
)

type SayImpl struct {
}

func (s *SayImpl) Hello(ctx context.Context, req *model.SayParam, rsp *model.SayResponse) error {
	fmt.Println("received", req.Msg)
	rsp.Header = make(map[string]*model.Pair)
	rsp.Header["name"] = &model.Pair{Key: 1, Values: "abc"}
	rsp.Msg = "hello world"
	rsp.Values = append(rsp.Values, "a", "b")
	rsp.Type = model.RespType_DESCEND
	return nil
}

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
	// 注册 Handler
	rpcapi.RegisterSayHandler(service.Server(), new(SayImpl))

	bk := service.Server().Options().Broker

	if err := bk.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}
	if err := bk.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}

	defer bk.Disconnect()

	sub, err := bk.Subscribe(topic, func(p broker.Event) error {
		fmt.Println("receive subscribe msg")
		fmt.Printf("[sub]:Received Body: %s,Header:%s", string(p.Message().Body), p.Message().Header)
		fmt.Println()
		return nil
	})

	if err != nil {
		fmt.Println("subscribe failed")
		return
	}

	defer sub.Unsubscribe()
	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
