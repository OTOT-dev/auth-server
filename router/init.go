package router

import (
	"auth-server/api"
	"auth-server/config"
	"auth-server/middleware"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var (
	apiUser api.UserApi
	apiAuth api.AuthApi

	validate = middleware.ValidatorMiddleware
	typeof   = reflect.TypeOf
)

func InitRouter() {
	engine := gin.New()
	engine.Use(gin.Recovery())

	userRouterGroup := engine.Group("/api/v1")
	userRouterGroup.Use(middleware.JWT())
	userRouter(userRouterGroup)
	port := config.ServerPort
	runParams := config.ServerHost + ":" + strconv.Itoa(port)
	log.Println("master server at ", runParams)
	if err := engine.Run(runParams); err != nil {
		log.Error(err)
		return
	}
}
