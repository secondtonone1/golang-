package authservice

import (
	// 导入生成的 protobuf 代码
	"context"
	"fmt"
	authpb "golang-/grpcservice/auth/authproto"
	dbpb "golang-/grpcservice/db/dbproto"
	config "golang-/grpcservice/serviceconfig"

	proto "github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
)

type AuthServiceImpl struct {
	userdata  map[string]int32
	accountid int32
	dbclient  dbpb.DBServiceClient
	dbconn    *grpc.ClientConn
}

func (ai *AuthServiceImpl) Auth(ctx context.Context, req *authpb.AuthReq) (*authpb.AuthRsp, error) {
	authname := req.GetName()
	value, ok := ai.userdata[authname]
	if !ok {
		return &authpb.AuthRsp{Errorid: config.RSP_ACTNOTREG}, nil
	}
	return &authpb.AuthRsp{Name: authname, Userid: value, Errorid: config.RSP_SUCCESS}, nil
}

func (ai *AuthServiceImpl) GenAccountid() int32 {
	ai.accountid++
	return ai.accountid
}

func (ai *AuthServiceImpl) AuthAdd(ctx context.Context, req *authpb.AuthReq) (*authpb.AuthRsp, error) {
	authname := req.GetName()
	_, ok := ai.userdata[authname]
	if ok {
		return &authpb.AuthRsp{Errorid: config.RSP_ACTHASREG}, nil
	}
	accountid := ai.GenAccountid()
	ai.userdata[authname] = accountid

	keystr := "account_" + authname
	pbstr, err := proto.Marshal(&dbpb.DBUsrData{
		Name:   authname,
		Userid: accountid,
	})

	if err != nil {
		return &authpb.AuthRsp{Errorid: config.RSP_PROTOMARSHERR}, nil
	}
	_, err = ai.dbclient.SaveData(ctx, &dbpb.DBSaveReq{Key: keystr, Value: string(pbstr)})

	if err != nil {
		return &authpb.AuthRsp{Errorid: config.RSP_SAVEMSGERR}, nil
	}

	genidbytes, err := proto.Marshal(&dbpb.DBGenuid{
		Genuid: accountid,
	})

	if err != nil {
		return &authpb.AuthRsp{Errorid: config.RSP_PROTOMARSHERR}, nil
	}

	_, err = ai.dbclient.SaveData(ctx, &dbpb.DBSaveReq{Key: "genuid_",
		Value: string(genidbytes)})

	if err != nil {
		return &authpb.AuthRsp{Errorid: config.RSP_SAVEMSGERR}, nil
	}

	return &authpb.AuthRsp{Name: authname, Userid: accountid, Errorid: config.RSP_SUCCESS}, nil
}

func (ai *AuthServiceImpl) Closervice() {
	if ai == nil {
		return
	}

	if ai.dbconn == nil {
		return
	}

	defer ai.dbconn.Close()
	if ai.userdata == nil {
		return
	}
	ai.userdata = nil

}

func NewAuthServiceImpl() (authpb.AuthServiceServer, error) {
	conn, err := grpc.Dial(config.DBaddress, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Did not connect: ", err)
		return nil, config.ErrAuthServerInit
	}

	client := dbpb.NewDBServiceClient(conn)
	dbrsp, err := client.LoadUsrData(context.Background(), &dbpb.DBUsrDataReq{})
	if err != nil {
		conn.Close()
		return nil, config.ErrAuthServerInit
	}
	ai := new(AuthServiceImpl)

	ai.dbclient = client
	ai.dbconn = conn
	ai.userdata = make(map[string]int32)
	for _, data := range dbrsp.GetUsrdatas() {
		ai.userdata[data.GetName()] = data.GetUserid()
	}
	ai.accountid = dbrsp.GetAccountid()
	return ai, nil
}
