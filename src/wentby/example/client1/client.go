package main

import (
	"time"
	"wentby/netmodel"
	"wentby/protocol"
)

func main() {

	cs, err := netmodel.Dial("tcp4", "127.0.0.1:10006")
	if err != nil {
		return
	}
	packet := new(protocol.MsgPacket)
	packet.Head.Id = 1
	packet.Head.Len = 5
	packet.Body.Data = []byte("Hello")
	cs.Send(packet)
	for {
		time.Sleep(time.Second * time.Duration(1))
	}
}
