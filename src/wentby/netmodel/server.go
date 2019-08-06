package netmodel

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"wentby/config"
)

func NewServer() (*WtServer, error) {
	address := config.SERVER_IP + ":" + strconv.Itoa(config.SERVER_PORT)
	listenert, err := net.Listen(config.SERVER_TYPE, address)
	if err != nil {
		fmt.Println("listen failed !!!")
		return nil, config.ErrListenFailed
	}

	return &WtServer{listener: listenert, stopedChan: make(chan struct{}), once: &sync.Once{}}, nil
}

type WtServer struct {
	listener   net.Listener
	stopedChan chan struct{}
	once       *sync.Once
}

func (wt *WtServer) Close() {
	wt.once.Do(func() {
		if wt.listener != nil {
			defer wt.listener.Close()
		}
		//send signal to all session
		close(wt.stopedChan)
	})

}

func (wt *WtServer) acceptLoop() error {
	tcpConn, err := wt.listener.Accept()
	if err != nil {
		fmt.Println("Accept error!")
		return config.ErrAcceptFailed
	}

	newsess := NewSession(tcpConn, wt.stopedChan)
	newsess.Start()
	fmt.Println("A client connected :" + tcpConn.RemoteAddr().String())
	return nil
}

func (wt *WtServer) AcceptLoop() {

	stopsignal := make(chan os.Signal) // 接收系统中断信号
	var shutdownSignals = []os.Signal{os.Interrupt, syscall.SIGTERM, syscall.SIGINT}
	signal.Notify(stopsignal, shutdownSignals...)

	for {
		select {
		case <-stopsignal:
			fmt.Println("server stop by signal")
			wt.Close()
			return
		default:
			wt.acceptLoop()
		}

	}

}
