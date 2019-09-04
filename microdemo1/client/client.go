package main

import (
	"context"
	"fmt"
	model "golang-/microdemo1/proto"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
)

func main() {
	// 我这里用的etcd 做为服务发现，如果使用consul可以去掉
	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"localhost:2379", "localhost:3379", "localhost:4379",
		}
	})

	// 初始化服务
	service := micro.NewService(
		micro.Registry(reg),
	)
	service.Init()
	sayClent := model.NewSayService("lp.srv.eg1", service.Client())

	rsp, err := sayClent.Hello(context.Background(), &model.SayParam{Msg: "hello server"})
	if err != nil {
		panic(err)
	}

	fmt.Println(rsp)

}
