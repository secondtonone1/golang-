package netmodel

import (
	"net"
	"sync/atomic"
	"wentby/protocol"
)

type Session struct {
	conn       net.Conn
	closed     int32
	stopedChan chan<- struct{}
	protocol   protocol.ProtocolInter
}

func NewSession(connt net.Conn, stopchan chan<- struct{}, pl protocol.ProtocolInter) *Session {
	sess := &Session{
		conn:       connt,
		closed:     -1,
		stopedChan: stopchan,
		protocol:   pl,
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
	}
	return nil
}

func (se *Session) sendLoop() {
	defer se.Close()

}

func (se *Session) recvLoop() {
	defer se.Close()
	var packet interface{}
	var err error
	for {
		packet, err := se.pl.ReadPacket(se.conn)
		if packet == nil or err == nil{
			return
		}

		//handle msg packet
	}
}
