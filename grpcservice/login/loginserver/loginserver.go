package loginserver

import (
	"fmt"
	//"log"
	"net"

	config "golang-/grpcservice/serviceconfig"

	loginpb "golang-/grpcservice/login/loginproto"
	loginservice "golang-/grpcservice/login/loginservice"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func LoginStart() {
	// 启动 gRPC 服务器。
	lis, err := net.Listen("tcp", config.Loginaddress)
	if err != nil {
		//log.Fatalf("failed to listen: %v", err)
		fmt.Printf("failed to listen: %v", err)
		fmt.Println(" ")
	}
	defer lis.Close()
	s := grpc.NewServer()

	loginserver := loginservice.NewLoginServiceImpl()
	if loginserver == nil {
		fmt.Println("login server create failed")
		return
	}
	defer loginserver.(*loginservice.LoginServiceImpl).Closervice()
	// 注册服务到 gRPC 服务器，会把已定义的 protobuf 与自动生成的代码接口进行绑定。
	loginpb.RegisterLoginServiceServer(s, loginserver)

	// 在 gRPC 服务器上注册 reflection 服务。
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		//log.Fatalf("failed to serve: %v", err)
		fmt.Printf("failed to serve: %v", err)
	}
}
