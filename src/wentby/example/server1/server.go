package main

import (
	"wentby/logic"
	"wentby/netmodel"
)

func main() {
	logic.RegHandlers()
	wt, err := netmodel.NewServer()
	if err != nil {
		return
	}
	defer wt.Close()
	wt.AcceptLoop()
}
