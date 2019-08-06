package wtwebsocket

import (
	"fmt"
	"time"

	"golang.org/x/net/websocket"
)

func RegWebServerHandlers() {
	//weblogic.RegHelloWorld()
	var svrConnHandler websocket.Handler = func(conn *websocket.Conn) {
		request := make([]byte, 128)
		defer conn.Close()
		for {
			readLen, err := conn.Read(request)
			if err != nil {
				break
			}

			//socket被关闭了
			if readLen == 0 {
				fmt.Println("Client connection close!")
				break
			} else {
				//输出接收到的信息
				fmt.Println(string(request[:readLen]))

				time.Sleep(time.Second)
				//发送
				conn.Write([]byte("World !"))
			}

			request = make([]byte, 128)
		}
	}

}
