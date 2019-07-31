package config

import "errors"

const (
	SERVER_IP   = "127.0.0.1"
	SERVER_PORT = 10006
	SERVER_TYPE = "tcp"
)

var (
	ErrListenFailed = errors.New("Listen Failed Error")
	ErrAcceptFailed = errors.New("Accept Failed Error")
)
