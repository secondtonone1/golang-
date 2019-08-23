package dbservice

import (

	// 导入生成的 protobuf 代码
	"fmt"
	dbpb "golang-/grpcservice/db/dbproto"
	"sync"

	config "golang-/grpcservice/serviceconfig"

	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
)

func NewDBServiceImpl() dbpb.DBServiceServer {
	return &DBServiceImpl{DBchan: make(chan *MsgPacket, config.SAVECHAN_SIZE),
		Savechan: make(chan int32, config.SAVEGOROUTINE_NUM),
		Once:     &sync.Once{},
	}
}

type MsgPacket struct {
	key   []byte
	value []byte
}

type DBServiceImpl struct {
	DBchan   chan *MsgPacket
	Once     *sync.Once
	Savechan chan int32
}

func (db *DBServiceImpl) StartSaveGoroutine() {
	for i := 0; i < config.SAVEGOROUTINE_NUM; i++ {
		go func(dbchan <-chan *MsgPacket, ip int32, savechan chan<- int32) {
			defer func(it int32) {
				fmt.Println("savegoroutine exited ,id is: ", it, " !")

			}(ip)
			fmt.Println("savegoroutine begined ,id is: ", ip, " !")
			for {
				select {
				case msg, ok := <-dbchan:
					fmt.Println("save msg", msg)
					if !ok {
						fmt.Println("dbmgr exited!")
						return
					}
					saverr := GetDBManagerIns().PutData(msg.key, msg.value)
					if saverr != nil {
						savechan <- ip
						return
					}
				}
			}

		}(db.DBchan, int32(i), db.Savechan)
	}
}

func (db *DBServiceImpl) PostMsgtoSave(msg *MsgPacket) error {
	if len(db.Savechan) >= config.SAVEGOROUTINE_NUM {
		fmt.Println("all save routines exit")
		return config.ErrAllSaveRoutinesClose
	}
	db.DBchan <- msg
	return nil
}

func (db *DBServiceImpl) Closervice() {
	db.Once.Do(func() {
		close(db.DBchan)
		if len(db.Savechan) >= config.SAVEGOROUTINE_NUM {
			close(db.Savechan)
		}
	})
}

func (db *DBServiceImpl) SaveData(ctx context.Context, req *dbpb.DBSaveReq) (*dbpb.DBSaveRsp, error) {
	err := db.PostMsgtoSave(&MsgPacket{key: []byte(req.Key), value: []byte(req.Value)})
	if err != nil {
		fmt.Println("db post msg failed")
		return &dbpb.DBSaveRsp{Errorid: config.RSP_SAVEMSGERR}, nil
	}

	return &dbpb.DBSaveRsp{Errorid: config.RSP_SUCCESS, Key: req.Key, Value: req.Value}, nil
}

func (db *DBServiceImpl) LoadUsrData(ctx context.Context, req *dbpb.DBUsrDataReq) (*dbpb.DBUsrDataRsp, error) {
	actdata := GetDBManagerIns().LoadAccountData()
	loadrsp := new(dbpb.DBUsrDataRsp)
	for _, data := range actdata {
		usrpb := &dbpb.DBUsrData{}
		proto.Unmarshal(data, usrpb)
		loadrsp.Usrdatas = append(loadrsp.Usrdatas, usrpb)
	}

	actid := GetDBManagerIns().LoadGenuid()
	actpb := &dbpb.DBGenuid{}
	proto.Unmarshal(actid, actpb)
	loadrsp.Accountid = actpb.GetGenuid()
	return loadrsp, nil
}
