package registerservice

import (

	// 导入生成的 protobuf 代码
	authpb "golang-/grpcservice/auth/authproto"
	registerpb "golang-/grpcservice/register/registerproto"
	config "golang-/grpcservice/serviceconfig"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func NewRegisterServiceImpl() registerpb.RegisterServiceServer {
	conn, err := grpc.Dial(config.Authaddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
		return nil
	}
	client := authpb.NewAuthServiceClient(conn)

	return &RegisterServiceImpl{authclient: client, authconn: conn}
}

type RegisterServiceImpl struct {
	authclient authpb.AuthServiceClient
	authconn   *grpc.ClientConn
}

func (rs *RegisterServiceImpl) Register(ctx context.Context, req *registerpb.RegisterReq) (*registerpb.RegisterRsp, error) {

	authrsp, err := rs.authclient.AuthAdd(ctx, &authpb.AuthReq{Name: req.GetName()})

	if err != nil {
		return nil, err
	}
	return &registerpb.RegisterRsp{Name: authrsp.GetName(), Userid: authrsp.GetUserid(), Errorid: authrsp.GetErrorid()}, nil

}

func (ls *RegisterServiceImpl) Closervice() {
	if ls == nil {
		return
	}
	if ls.authclient == nil {
		return
	}
	defer ls.authconn.Close()
}
