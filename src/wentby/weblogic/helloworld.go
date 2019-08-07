package weblogic

import (
	"fmt"
	"net/http"
	"wentby/config"

	"golang.org/x/net/websocket"
)

func RegHelloWorld(pattern string) {

	var svrConnHandler websocket.Handler = func(conn *websocket.Conn) {
		request := make([]byte, config.MAXMESSAGE_LEN)

		readLen, err := conn.Read(request)
		if err != nil {
			fmt.Println(config.ErrWebSocketRead.Error())
			return
		}

		//socket被关闭了
		if readLen == 0 {
			fmt.Println(config.ErrWebSocketClosed.Error())
			return
		}

		fmt.Println(string(request[:readLen]))
		conn.Write([]byte("Recieve Hello World Msg!"))
	}

	http.Handle(pattern, svrConnHandler)
}
