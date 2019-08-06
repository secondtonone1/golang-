package main

import (
	"fmt"
	"time"
	"wentby/netmodel"
	"wentby/protocol"
)

func main() {
	for i := 0; i < 100; i++ {
		go func() {
			cs, err := netmodel.Dial("tcp4", "127.0.0.1:10006")
			if err != nil {
				return
			}
			var i int16
			for i = 0; i < 100; i++ {
				packet := new(protocol.MsgPacket)
				packet.Head.Id = 1
				packet.Head.Len = 5
				packet.Body.Data = []byte("Hello")
				cs.Send(packet)
				packetrsp, err := cs.Recv()
				if err != nil {
					fmt.Println("receive error")
					return
				}

				datarsp := packetrsp.(*protocol.MsgPacket)
				fmt.Println("packet id is", datarsp.Head.Id)
				fmt.Println("packet len is", datarsp.Head.Len)
				fmt.Println("packet data is", string(datarsp.Body.Data))
				time.Sleep(time.Millisecond * time.Duration(10))
			}
			fmt.Println("circle times are ", i)
		}()

		time.Sleep(time.Second * time.Duration(1))
	}
	for {
		time.Sleep(time.Second * time.Duration(15))
	}
}
