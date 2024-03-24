package config

import (
	"os"
	"strconv"
	"time"
)

var (
	ServerName    = "auth-server" // 服务名称
	ServerHost    string
	ServerPort    int
	DebugMode     bool
	JwtSecret     []byte          //jwt密钥
	JwtExpire     = 3 * time.Hour //jwt过期时间
	SessionExpire = 3600 * 1      // session过期时间 1h
	SessionSecret = "bGjW7xiMrxC9lmXN"
	DataBaseHost  string
	DataBasePort  int
	DataBaseName  string
	LogPath       = "./log"
)

func initConfig() {
	//环境变量生效优先级 命令行》环境变量》配置文件
	if serverHost := os.Getenv("SERVER_HOST"); serverHost != "" {
		ServerHost = serverHost
	}

	if serverPort := os.Getenv("SERVER_PORT"); serverPort != "" {
		ServerPort, _ = strconv.Atoi(serverPort)
	}
	if debugMode := os.Getenv("DEBUG_MODE"); debugMode != "" {
		DebugMode, _ = strconv.ParseBool(debugMode)
	}
	if jwtSecret := os.Getenv("JWT_SECRET"); jwtSecret != "" {
		JwtSecret = []byte(jwtSecret)
	}
	if dataBaseHost := os.Getenv("DATABASE_HOST"); dataBaseHost != "" {
		DataBaseHost = dataBaseHost
	}
	if dataBasePort := os.Getenv("DATABASE_PORT"); dataBasePort != "" {
		DataBasePort, _ = strconv.Atoi(dataBasePort)
	}
	if dataBaseName := os.Getenv("DATABASE_NAME"); dataBaseName != "" {
		DataBaseName = dataBaseName
	}
	if sessionSecret := os.Getenv("SESSION_SECRET"); sessionSecret != "" {
		SessionSecret = sessionSecret
	}
	if logPath := os.Getenv("LOG_PATH"); logPath != "" {
		LogPath = logPath
	}
}
