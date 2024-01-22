package router

import (
	"auth-server/api"
	"auth-server/config"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"strconv"
)

var (
	apiUser api.UserApi
)

func InitRouter() {
	engine := gin.New()
	engine.Use(gin.Recovery())

	userRouterGroup := engine.Group("/api/v1")
	userRouter(userRouterGroup)
	port := config.ServerPort
	runParams := config.ServerHost + ":" + strconv.Itoa(port)
	log.Println("master server at ", runParams)
	if err := engine.Run(runParams); err != nil {
		log.Error(err)
		return
	}
}
