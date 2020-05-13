package main

import (
	components "golang-/seckill/seclogic/components"
	_ "golang-/seckill/seclogic/service"
)

func main() {
	defer components.ReleaseRsc()
}
