package main

import (
	"fmt"
	"wentby/config"

	"golang.org/x/net/websocket"
)

var url = "ws://127.0.0.1:9998/"
var origin = "http://127.0.0.1:9998/"

func main() {
	conn, err := websocket.Dial(url, "", origin)
	if err != nil {
		fmt.Println(config.ErrWebSocketDail.Error())
	}

	_, err = conn.Write([]byte("Hello !"))
	if err != nil {
		fmt.Println(config.ErrWebSocketWrite)
		return
	}
	readdata := make([]byte, 1024)
	readlen, err := conn.Read(readdata)
	if err != nil {
		fmt.Println(config.ErrWebSocketRead.Error())
		return
	}

	fmt.Println("client recieve msg is : ", string(readdata[:readlen]))
}
