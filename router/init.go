package router

import (
	"auth-server/api"
	"auth-server/config"
	"auth-server/middleware"
	"strconv"

	"github.com/gin-gonic/contrib/sessions"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var (
	apiUser api.UserApi
	apiAuth api.AuthApi
)

func InitRouter() {
	engine := gin.New()
	engine.Use(gin.Recovery())

	// session 设置
	store := sessions.NewCookieStore([]byte(config.SessionSecret))
	store.Options(sessions.Options{
		Path:   "/",
		MaxAge: config.SessionExpire, // 设置超时时间为一个小时
	})
	engine.Use(sessions.Sessions("sid", store))

	// 登陆认证相关路由
	authRouterGroup := engine.Group("/auth")
	authRouter(authRouterGroup)
	// 用户登陆
	userRouterGroup := engine.Group("/api/v1")
	userRouterGroup.Use(middleware.SessionAuth())

	userRouter(userRouterGroup)
	port := config.ServerPort
	runParams := config.ServerHost + ":" + strconv.Itoa(port)
	log.Println("master server at ", runParams)
	if err := engine.Run(runParams); err != nil {
		log.Error(err)
		return
	}
}
