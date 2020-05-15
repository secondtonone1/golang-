package main

import (
	components "golang-/seckill/seclogic/components"
	service "golang-/seckill/seclogic/service"
)

func main() {
	defer components.ReleaseRsc()
	service.InitRedisRWGoroutine()
}
