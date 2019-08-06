package logic

import (
	"fmt"
	"wentby/config"
	"wentby/netmodel"
	"wentby/protocol"
)

/*
func HelloworldReq(se interface{}, param interface{}) error {
	msgpacket, ok := param.(*protocol.MsgPacket)
	if !ok {
		return config.ErrTypeAssertain
	}

	session, ok := se.(*netmodel.Session)
	fmt.Println("Server recieve from ", session.RawConn().RemoteAddr().String())
	fmt.Println("Server Recv Msg is ", string(msgpacket.Body.Data))
	helloworldrsp := new(protocol.MsgPacket)
	helloworldrsp.Head.Id = HELLOWORLD_RSP
	helloworldrsp.Head.Len = uint16(len("server recive msg hello world!"))
	helloworldrsp.Body.Data = []byte("server recive msg hello world!")
	err := session.AsyncSend(helloworldrsp)
	if err != nil {
		fmt.Println("Handle Msg HelloworldReq failed")
		return config.ErrHelloWorldReqFailed
	}
	return nil
}

func HelloworldRsp(session interface{}, param interface{}) error {
	return nil
}
*/
func RegHelloWorld() {
	var HelloworldReq netmodel.CallBackFunc = func(se interface{}, param interface{}) error {
		msgpacket, ok := param.(*protocol.MsgPacket)
		if !ok {
			return config.ErrTypeAssertain
		}

		session, ok := se.(*netmodel.Session)
		fmt.Println("Server recieve from ", session.RawConn().RemoteAddr().String())
		fmt.Println("Server Recv Msg is ", string(msgpacket.Body.Data))
		helloworldrsp := new(protocol.MsgPacket)
		helloworldrsp.Head.Id = HELLOWORLD_RSP
		helloworldrsp.Head.Len = uint16(len("server recive msg hello world!"))
		helloworldrsp.Body.Data = []byte("server recive msg hello world!")
		err := session.AsyncSend(helloworldrsp)
		if err != nil {
			fmt.Println("Handle Msg HelloworldReq failed")
			return config.ErrHelloWorldReqFailed
		}
		return nil
	}

	netmodel.MsgHandler.RegMsgHandler(HELLOWORLD_REQ, HelloworldReq)
}
