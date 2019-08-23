package main

import (
	"context"
	"fmt"
	"log"

	config "golang-/grpcservice/serviceconfig"

	dbpb "golang-/grpcservice/db/dbproto"

	"google.golang.org/grpc"
)

func main() {

	// 创建和服务器的一个连接
	conn, err := grpc.Dial(config.DBaddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	client := dbpb.NewDBServiceClient(conn)
	dbrsp, err := client.LoadUsrData(context.Background(), &dbpb.DBUsrDataReq{})

	if err != nil {
		fmt.Println("db failed , error is ", err.Error())
		return
	}

	if dbrsp.Errorid == config.RSP_SUCCESS {
		fmt.Println("db success")
		return
	}

}
