package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

//发送信息
func sender(conn net.Conn) {

	words := "Hello Server!"
	conn.Write([]byte(words))
	fmt.Println("send over")
	//接收服务端反馈
	buffer := make([]byte, 2048)
	var index int = 1
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			Log(conn.RemoteAddr().String(), "waiting server back msg error: ", err)
			return
		}
		Log(conn.RemoteAddr().String(), "receive server back msg: ", string(buffer[:n]))
		//conn.Write([]byte(buffer[:n]))
		data := "Hello Client Send Msg " + strconv.Itoa(index) + " times"
		conn.Write([]byte(data))
	}
}

//日志
func Log(v ...interface{}) {
	log.Println(v...)
}

func main() {
	server := "127.0.0.1:1024"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	fmt.Println("connection success")
	sender(conn)
}
