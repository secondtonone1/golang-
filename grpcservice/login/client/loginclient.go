package main

import (
	"context"
	"fmt"
	"log"

	config "golang-/grpcservice/serviceconfig"

	loginpb "golang-/grpcservice/login/loginproto"

	registerpb "golang-/grpcservice/register/registerproto"

	"google.golang.org/grpc"
)

func main() {

	// 创建和服务器的一个连接
	conn, err := grpc.Dial(config.Loginaddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
		return
	}
	defer conn.Close()
	client := loginpb.NewLoginServiceClient(conn)

	loginrsp, err := client.Login(context.Background(), &loginpb.LoginReq{Name: "Zack"})

	if err != nil {
		fmt.Println("grpc calls error!")
		return
	}

	if loginrsp.Errorid == config.RSP_SUCCESS {
		fmt.Println("login success")
		fmt.Println("loginrsp name is", loginrsp.Name)
		fmt.Println("loginrsp usrid is", loginrsp.Userid)
		fmt.Println("loginrsp errorid is", loginrsp.Errorid)
		return
	}

	if loginrsp.Errorid == config.RSP_ACTNOTREG {
		regcon, reger := grpc.Dial(config.Registeraddress, grpc.WithInsecure())
		if reger != nil {
			fmt.Println("Did not connect", err.Error())
			return
		}

		defer regcon.Close()
		regclient := registerpb.NewRegisterServiceClient(regcon)
		regrsp, reger := regclient.Register(context.Background(), &registerpb.RegisterReq{Name: "Zack"})
		if reger == nil {
			fmt.Println("register success")
			fmt.Println("registerrsp name is", regrsp.Name)
			fmt.Println("registerrsp usrid is", regrsp.Userid)
			fmt.Println("registerrsp errorid is", regrsp.Errorid)
			return
		}

		if regrsp.Errorid == config.RSP_ACTHASREG {
			fmt.Println("account has been reged")
			return
		}
		fmt.Println("Unkown register error!")
		return
	}

}
