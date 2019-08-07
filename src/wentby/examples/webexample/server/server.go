package main

import "wentby/wtwebsocket"

func main() {
	webserver := wtwebsocket.NewWtWebServer()
	webserver.Start()
}
