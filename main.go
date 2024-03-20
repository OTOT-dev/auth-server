package main

import (
	"auth-server/common"
	"auth-server/config"
	"auth-server/router"
)

func main() {
	//初始化日志
	common.InitLog(config.LogPath, config.ServerName)
	router.InitRouter()
}
