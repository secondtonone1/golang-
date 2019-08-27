package authserver

import (
	"fmt"
	"log"
	"net"

	config "golang-/grpcservice/serviceconfig"

	authpb "golang-/grpcservice/auth/authproto"
	authservice "golang-/grpcservice/auth/authservice"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func AuthStart() {
	// 启动 gRPC 服务器。
	lis, err := net.Listen("tcp", config.Authaddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()
	s := grpc.NewServer()

	// 注册服务到 gRPC 服务器，会把已定义的 protobuf 与自动生成的代码接口进行绑定。
	authserver, err := authservice.NewAuthServiceImpl()
	if err != nil {
		fmt.Println("authserver create failed")
		return
	}

	defer authserver.(*authservice.AuthServiceImpl).Closervice()
	authpb.RegisterAuthServiceServer(s, authserver)

	// 在 gRPC 服务器上注册 reflection 服务。
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
