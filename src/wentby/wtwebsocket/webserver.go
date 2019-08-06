package wtwebsocket

import (
	"net/http"
	"strconv"
	"wentby/config"

	"golang.org/x/net/websocket"
)

type WtWebServer struct {
}

func (wb *WtWebServer) RegWebHandler(pattern string, handler websocket.Handler) {
	http.Handle(pattern, handler)
}

func (wb *WtWebServer) ListenAndServe() error {
	address := config.SERVER_IP + ":" + strconv.Itoa(config.WEBSERVER_PORT)
	err := http.ListenAndServe(address, nil)
	return err
}
