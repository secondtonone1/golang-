package main

import (
	"context"
	"fmt"
	model "golang-/microdemo4/proto/model"
	"golang-/microdemo4/proto/rpcapi"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/broker/grpc"
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

	//订阅

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

	// 初始化服务
	service := micro.NewService(
		micro.Name("lp.srv.eg1"),
		micro.Registry(reg),
		micro.Broker(bk),
	)
	service.Init()
	// 注册 Handler
	rpcapi.RegisterSayHandler(service.Server(), new(SayImpl))
	// run server
	if err := service.Run(); err != nil {
		panic(err)
	}
}
