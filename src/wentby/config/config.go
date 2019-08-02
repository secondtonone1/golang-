package config

import (
	"errors"
)

const (
	SERVER_IP     = "127.0.0.1"
	SERVER_PORT   = 10006
	SERVER_TYPE   = "tcp"
	SENDCHAN_SIZE = 1024
)

var (
	ErrListenFailed  = errors.New("Listen Failed Error")
	ErrAcceptFailed  = errors.New("Accept Failed Error")
	ErrBuffOverflow  = errors.New("Buff Overflow Error")
	ErrBuffLenLess   = errors.New("Buff Length is not enough")
	ErrParaseMsgHead = errors.New("Parase Msg Head Failed")
	ErrTypeAssertain = errors.New("Type Assertain failed")
	ErrMsgLenLarge   = errors.New("Msg Length is too large")
	ErrReadAtLeast   = errors.New("Read at least error!")
	ErrMsgHandlerReg = errors.New("Msg Handler function not reg")
	ErrParamCallBack = errors.New("Param is not call back")
)
