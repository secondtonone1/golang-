package weblogic

import (
	"net/http"

	"golang.org/x/net/websocket"
)

func RegHelloWorld(pattern string, handler websocket.Handler) {
	http.Handle(pattern, handler)
}
