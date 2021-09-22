package main

import (
	_ "gin_template/module/pong"
	"gin_template/server"
	"gin_template/utils"
)

func init() {
	utils.WriteLogToFS()
}

func main() {
	server.Init()

	server.StartService()

	server.Run()
}
