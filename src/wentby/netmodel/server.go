package netmodel

import (
	"fmt"
	"net"
	"strconv"
	"wentby/config"
	"wentby/protocol"
)

func NewServer() (*WtServer, error) {
	address := config.SERVER_IP + ":" + strconv.Itoa(config.SERVER_PORT)
	listenert, err := net.Listen(config.SERVER_TYPE, address)
	if err != nil {
		fmt.Println("listen failed !!!")
		return nil, config.ErrListenFailed
	}

	return &WtServer{listener: listenert, stopedChan: make(chan struct{})}, nil
}

type WtServer struct {
	listener   net.Listener
	stopedChan chan struct{}
}

func (wt *WtServer) Close() {
	if wt.listener != nil {
		defer wt.listener.Close()
	}
	//send signal to all session
	close(wt.stopedChan)
}

func (wt *WtServer) AcceptLoop() {
	for {
		tcpConn, err := wt.listener.Accept()
		if err != nil {
			fmt.Println("Accept error!")
			continue
		}
		newsess := NewSession(tcpConn, wt.stopedChan, new(protocol.ProtocolImpl))
		newsess.Start()
		fmt.Println("A client connected :" + tcpConn.RemoteAddr().String())
	}
}
