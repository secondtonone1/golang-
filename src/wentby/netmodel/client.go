package netmodel

import (
	"net"
	"os"
	"os/signal"
	"syscall"
	"wentby/config"
	"wentby/protocol"
)

type ClientSess struct {
	conn       net.Conn
	stopedChan <-chan os.Signal
	protocol   protocol.ProtocolInter
}

func NewClientSess(connt net.Conn, stopchan <-chan os.Signal) *ClientSess {
	clss := &ClientSess{
		conn:       connt,
		stopedChan: stopchan,
		protocol:   new(protocol.ProtocolImpl),
	}
	tcpConn := clss.conn.(*net.TCPConn)
	tcpConn.SetNoDelay(true)
	tcpConn.SetReadBuffer(64 * 1024)
	tcpConn.SetWriteBuffer(64 * 1024)
	return clss
}

func Dial(network, address string) (*ClientSess, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	stopsignal := make(chan os.Signal) // 接收系统中断信号
	var shutdownSignals = []os.Signal{os.Interrupt, syscall.SIGTERM, syscall.SIGINT}
	signal.Notify(stopsignal, shutdownSignals...)
	return NewClientSess(conn, stopsignal), nil
}

func (cs *ClientSess) Send(packet interface{}) error {
	select {
	case <-cs.stopedChan:
		return config.ErrSignalStopped
	default:
		err := cs.protocol.WritePacket(cs.conn, packet)
		if err != nil {
			return config.ErrWritePacketFailed
		}

	}
	return nil
}

func (cs *ClientSess) Recv() (interface{}, error) {
	packet, err := cs.protocol.ReadPacket(cs.conn)
	if err != nil {
		return nil, config.ErrReadPacketFailed
	}

	return packet, nil
}
