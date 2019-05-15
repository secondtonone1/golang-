package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	netListen, err := net.Listen("tcp", "localhost:1024")
	CheckError(err)
	defer netListen.Close()
	Log("Waiting for clients ...")

	for {
		conn, err := netListen.Accept()
		if err != nil {
			fmt.Println("tcp error")
			continue
		}

		Log(conn.RemoteAddr().String(), " tcp connect success")
		go handleConnection(conn)
	}
}

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error %s", err.Error)
	}
}

//日志处理
func Log(v ...interface{}) {
	log.Println(v...)
}

func handleConnection(conn net.Conn) {
	buffer := make([]byte, 2048)
	for {
		//读取客户端传来的内容
		n, err := conn.Read(buffer)
		if err != nil {
			Log(conn.RemoteAddr().String(), "connection error: ", err)
			//当远程客户端连接发生错误（断开）后，终止此协程。
			return
		}
		Log(conn.RemoteAddr().String(), "receive data string:\n", string(buffer[:n]))
		strTemp := "Server got msg \" " + string(buffer[:n]) + "\" at " + time.Now().String()
		conn.Write([]byte(strTemp))
	}
}
