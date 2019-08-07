package netmodel

import (
	"fmt"
	"net"
	"sync"
	"sync/atomic"
	"time"
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
	lock       sync.Mutex
}

func NewSession(connt net.Conn, stopchan <-chan struct{}) *Session {
	sess := &Session{
		conn:       connt,
		closed:     -1,
		stopedChan: stopchan,
		protocol:   new(protocol.ProtocolImpl),
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
		close(se.asyncStop)
		close(se.sendChan)
	}
	return nil
}

//set read time out
//if u don't need to set read deadline, please not use it
func (se *Session) SetReadDeadline(delt time.Duration) {
	se.lock.Lock()
	se.conn.SetReadDeadline(time.Now().Add(delt)) // timeout
	defer se.lock.Unlock()
}

func (se *Session) sendLoop() {
	defer se.Close()
	defer func() {
		fmt.Println("send goroutine exit!")
	}()
	for {
		select {
		case <-se.stopedChan:
			return
		case <-se.asyncStop:
			return
		case packet, ok := <-se.sendChan:
			{
				if !ok {
					return
				}
				err := se.protocol.WritePacket(se.conn, packet)
				if err != nil {
					return
				}
			}
		}
	}
}

func (se *Session) recvLoop() {
	defer se.Close()
	defer func() {
		fmt.Println("recv goroutine exit!")
	}()
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
				//set read time out
				//se.SetReadDeadline(time.Minute)
				packet, err = se.protocol.ReadPacket(se.conn)
				if packet == nil || err != nil {
					fmt.Println("Read packet error ", err.Error())
					return
				}

				//handle msg packet
				hdres := MsgHandler.HandleMsgPacket(packet, se)
				if hdres != nil {
					fmt.Println(hdres.Error())
					return
				}
			}

		}

	}
}

func (se *Session) AsyncSend(packet interface{}) error {
	select {
	case <-se.asyncStop:
		return config.ErrAsyncSendStop
	case <-se.stopedChan:
		return config.ErrAsyncSendStop
	default:
		se.sendChan <- packet
		return nil
	}
}
