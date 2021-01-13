package main

import (
	"aurora/blog/api/common"
	"aurora/blog/api/server"
)

func main() {
	common.SetupConfig()
	server.InitServer()
}
