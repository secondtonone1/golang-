package main

import "wentby/webserver"

func main() {
	webserver := webserver.NewWtWebServer()
	webserver.Start()
}
