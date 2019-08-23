package main

import (
	"context"
	"fmt"
	"log"

	authpb "golang-/grpcservice/auth/authproto"
	config "golang-/grpcservice/serviceconfig"

	"google.golang.org/grpc"
)

func main() {
	// 创建和服务器的一个连接
	conn, err := grpc.Dial(config.Authaddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	client := authpb.NewAuthServiceClient(conn)

	authrsp, err := client.Auth(context.Background(), &authpb.AuthReq{Name: "Zack"})

	if err != nil {
		fmt.Println("auth failed , error is ", err.Error())
		return
	}

	if authrsp.Errorid != config.RSP_SUCCESS {
		return
	}
	fmt.Println("auth success")
	fmt.Println("authrsp name is", authrsp.Name)
	fmt.Println("authrsp usrid is", authrsp.Userid)
	fmt.Println("authrsp errorid is", authrsp.Errorid)
}
