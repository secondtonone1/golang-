package netmodel

import (
	"net"
	"sync/atomic"
	"wentby/config"
	"wentby/protocol"
)

type Session struct {
	conn       net.Conn
	closed     int32
	stopedChan <-chan struct{}
	protocol   protocol.ProtocolInter
	asyncStop  chan struct{}
	sendChan   chan interface{}
}

func NewSession(connt net.Conn, stopchan <-chan struct{}, pl protocol.ProtocolInter) *Session {
	sess := &Session{
		conn:       connt,
		closed:     -1,
		stopedChan: stopchan,
		protocol:   pl,
		sendChan:   make(chan interface{}, config.SENDCHAN_SIZE),
		asyncStop:  make(chan struct{}),
	}
	tcpConn := sess.conn.(*net.TCPConn)
	tcpConn.SetNoDelay(true)
	tcpConn.SetReadBuffer(64 * 1024)
	tcpConn.SetWriteBuffer(64 * 1024)
	return sess
}

func (se *Session) RawConn() *net.TCPConn {
	return se.conn.(*net.TCPConn)
}

func (se *Session) Start() {
	if atomic.CompareAndSwapInt32(&se.closed, -1, 0) {
		go se.sendLoop()
		go se.recvLoop()
	}
}

// Close the session, destory other resource.
func (se *Session) Close() error {
	if atomic.CompareAndSwapInt32(&se.closed, 0, 1) {
		se.conn.Close()
		select {
		case <-se.asyncStop:
			return nil
		default:
			close(se.asyncStop)
			close(se.sendChan)
		}
	}
	return nil
}

func (se *Session) sendLoop() {
	defer se.Close()

	for {
		select {
		case <-se.stopedChan:
			return
		case <-se.asyncStop:
			return
		default:
			{
				//packet:=<-se.sendChan
				
			}
		}
	}

}

func (se *Session) recvLoop() {
	defer se.Close()

	var packet interface{}
	var err error
	for {

		select {
		case <-se.stopedChan:
			return
		case <-se.asyncStop:
			return
		default:
			{
				packet, err = se.protocol.ReadPacket(se.conn)
				if packet == nil || err != nil {
					return
				}

				//handle msg packet
				hdres := MsgHandler.HandleMsgPacket(packet, se)
				if hdres != nil {
					return
				}
			}

		}

	}
}

func (se *Session) asyncSend(packet interface{}) error{
	select {
	case <- se.asyncStop:
		return config.ErrAsyncSendStop
	case <- se.stopedChan:
		return config.ErrAsyncSendStop
	default:
		se.sendChan <- packet
		return nil
	}
}
