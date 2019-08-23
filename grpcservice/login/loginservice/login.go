package loginservice

import (

	// 导入生成的 protobuf 代码
	authpb "golang-/grpcservice/auth/authproto"
	loginpb "golang-/grpcservice/login/loginproto"
	config "golang-/grpcservice/serviceconfig"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func NewLoginServiceImpl() loginpb.LoginServiceServer {
	conn, err := grpc.Dial(config.Authaddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
		return nil
	}
	client := authpb.NewAuthServiceClient(conn)

	return &LoginServiceImpl{authclient: client, authconn: conn}
}

type LoginServiceImpl struct {
	authclient authpb.AuthServiceClient
	authconn   *grpc.ClientConn
}

func (ls *LoginServiceImpl) Login(ctx context.Context, req *loginpb.LoginReq) (*loginpb.LoginRsp, error) {
	loginame := req.GetName()

	authrsp, err := ls.authclient.Auth(ctx, &authpb.AuthReq{Name: loginame})
	if err != nil {
		return nil, err
	}

	return &loginpb.LoginRsp{
		Errorid: authrsp.Errorid,
		Name:    loginame,
		Userid:  authrsp.Userid,
	}, nil

}

func (ls *LoginServiceImpl) Closervice() {
	if ls == nil {
		return
	}
	if ls.authclient == nil {
		return
	}
	defer ls.authconn.Close()
}
