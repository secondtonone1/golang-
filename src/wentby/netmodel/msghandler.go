package netmodel

import (
	"wentby/config"
	"wentby/protocol"
)

type MsgHandlerInter interface {
	HandleMsgPacket(param interface{}) error
	RegMsgHandler(param interface{}) error
}

type CallBackFunc func(session interface{}, param interface{}) error
type MsgHandlerImpl struct {
	cbfuncs map[uint16]CallBackFunc
}

func (mh *MsgHandlerImpl) HandleMsgPacket(param interface{}, se interface{}) error {

	var (
		msgpacket *protocol.MsgPacket
		callback  CallBackFunc
		ok        bool
		session   *Session
	)
	if msgpacket, ok = param.(*protocol.MsgPacket); !ok {
		return config.ErrTypeAssertain
	}

	if session, ok = se.(*Session); !ok {
		return config.ErrTypeAssertain
	}

	if callback, ok = mh.cbfuncs[msgpacket.Head.Id]; !ok {
		//不存在
		return config.ErrMsgHandlerReg
	}

	return callback(session, param)
}

func (mh *MsgHandlerImpl) RegMsgHandler(cbid uint16, param interface{}) error {
	var (
		callback CallBackFunc
		ok       bool
	)

	if callback, ok = param.(CallBackFunc); !ok {
		return config.ErrParamCallBack
	}

	mh.cbfuncs[cbid] = callback
	return nil
}

var MsgHandler = &MsgHandlerImpl{cbfuncs: make(map[uint16]CallBackFunc)}
