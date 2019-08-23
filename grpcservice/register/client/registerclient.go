package main

import (
	"context"
	"fmt"
	"log"

	config "golang-/grpcservice/serviceconfig"

	reginpb "golang-/grpcservice/register/registerproto"

	"google.golang.org/grpc"
)

func main() {

	// 创建和服务器的一个连接
	conn, err := grpc.Dial(config.Registeraddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	client := reginpb.NewRegisterServiceClient(conn)
	regrsp, err := client.Register(context.Background(), &reginpb.RegisterReq{Name: "Zack"})

	if err != nil {
		fmt.Println("login failed , error is ", err.Error())
		return
	}

	if regrsp.Errorid == config.RSP_SUCCESS {
		fmt.Println("login success")
		fmt.Println("loginrsp name is", regrsp.Name)
		fmt.Println("loginrsp usrid is", regrsp.Userid)
		fmt.Println("loginrsp errorid is", regrsp.Errorid)
		return
	}

}
