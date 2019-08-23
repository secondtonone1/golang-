package main

import (
	"fmt"
	"log"
	"net"

	config "golang-/grpcservice/serviceconfig"

	registerpb "golang-/grpcservice/register/registerproto"
	registerservice "golang-/grpcservice/register/registerservice"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// 启动 gRPC 服务器。
	lis, err := net.Listen("tcp", config.Registeraddress)
	if err != nil {
		//log.Fatalf("failed to listen: %v", err)
		fmt.Printf("failed to listen: %v", err)
		fmt.Println(" ")
	}
	defer lis.Close()
	s := grpc.NewServer()

	registerserver := registerservice.NewRegisterServiceImpl()
	if registerserver == nil {
		fmt.Println("login server create failed")
		return
	}
	defer registerserver.(*registerservice.RegisterServiceImpl).Closervice()
	// 注册服务到 gRPC 服务器，会把已定义的 protobuf 与自动生成的代码接口进行绑定。
	registerpb.RegisterRegisterServiceServer(s, registerserver)

	// 在 gRPC 服务器上注册 reflection 服务。
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
