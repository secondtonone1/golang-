package wtwebsocket

import (
	"fmt"
	"net/http"
	"strconv"
	"wentby/config"
	"wentby/weblogic"
)

type WtWebServer struct {
}

func (wb *WtWebServer) RegWebHandler() {
	weblogic.RegWebServerHandlers()
}

func (wb *WtWebServer) ListenAndServe() error {
	address := config.SERVER_IP + ":" + strconv.Itoa(config.WEBSERVER_PORT)
	err := http.ListenAndServe(address, nil)
	return err
}

func (wb *WtWebServer) Start() {
	wb.RegWebHandler()
	err := wb.ListenAndServe()
	if err != nil {
		fmt.Println(config.ErrWebListenFailed.Error())
		return
	}
}

func NewWtWebServer() *WtWebServer {
	return &WtWebServer{}
}
