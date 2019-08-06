package netmodel

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"wentby/config"

	"golang.org/x/net/websocket"
)

//错误处理函数
func checkErr(err error, extra string) bool {
	if err != nil {
		formatStr := " Err : %s\n"
		if extra != "" {
			formatStr = extra + formatStr
		}

		fmt.Fprintf(os.Stderr, formatStr, err.Error())
		return true
	}

	return false
}

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
