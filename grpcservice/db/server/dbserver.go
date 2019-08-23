package main

import (
	"fmt"
	"log"
	"net"

	config "golang-/grpcservice/serviceconfig"

	dbpb "golang-/grpcservice/db/dbproto"
	dbservice "golang-/grpcservice/db/dbservice"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// 启动 gRPC 服务器。
	lis, err := net.Listen("tcp", config.DBaddress)
	if err != nil {
		//log.Fatalf("failed to listen: %v", err)
		fmt.Printf("failed to listen: %v", err)
		fmt.Println(" ")
	}
	defer lis.Close()
	s := grpc.NewServer()

	dbserver := dbservice.NewDBServiceImpl()

	if dbserver == nil {
		fmt.Println("db server create failed")
		return
	}
	dbserver.(*dbservice.DBServiceImpl).StartSaveGoroutine()
	defer dbserver.(*dbservice.DBServiceImpl).Closervice()

	dbmgr := dbservice.GetDBManagerIns()
	if dbmgr == nil {
		fmt.Println("db manager create failed")
		return
	}

	err = dbmgr.InitDB("./lvdb")
	if err != nil {
		fmt.Println("db manager init failed")
		return
	}

	defer dbmgr.CloseDB()

	// 注册服务到 gRPC 服务器，会把已定义的 protobuf 与自动生成的代码接口进行绑定。
	dbpb.RegisterDBServiceServer(s, dbserver)

	// 在 gRPC 服务器上注册 reflection 服务。
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
