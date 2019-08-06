package config

import (
	"errors"
)

const (
	SERVER_IP      = "127.0.0.1"
	SERVER_PORT    = 10006
	SERVER_TYPE    = "tcp"
	SENDCHAN_SIZE  = 1024
	MAXMESSAGE_LEN = 1024
	WEBSERVER_PORT = 9998
)

var (
	ErrListenFailed        = errors.New("Listen Failed Error")
	ErrAcceptFailed        = errors.New("Accept Failed Error")
	ErrBuffOverflow        = errors.New("Buff Overflow Error")
	ErrBuffLenLess         = errors.New("Buff Length is not enough")
	ErrParaseMsgHead       = errors.New("Parase Msg Head Failed")
	ErrTypeAssertain       = errors.New("Type Assertain failed")
	ErrMsgLenLarge         = errors.New("Msg Length is too large")
	ErrReadAtLeast         = errors.New("Read at least error!")
	ErrMsgHandlerReg       = errors.New("Msg Handler function not reg")
	ErrParamCallBack       = errors.New("Param is not call back")
	ErrAsyncSendStop       = errors.New("async send chan is stopped")
	ErrSessChanStoped      = errors.New(" session chan is stopped")
	ErrPacketEmpty         = errors.New("Packet is empty")
	ErrWritePacketFailed   = errors.New("Write packet failed")
	ErrConnWriteFailed     = errors.New("Connection Wrtie Failed")
	ErrSignalStopped       = errors.New("Signal Stopped")
	ErrReadPacketFailed    = errors.New("read packet failed!")
	ErrHelloWorldReqFailed = errors.New("Handle Msg Hellow World Req Failed")
)
