package router

import (
	"auth-server/api"
	"auth-server/config"
	"auth-server/docs"
	"auth-server/middleware"
	"strconv"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/contrib/sessions"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var (
	apiAuth api.AuthApi
	apiUser api.UserApi
)

var (
	authRoutesPrefix    = "/auth"
	serviceRoutesPrefix = "/api/v1"
)

var sessionName = "sid"

func InitRouter() {
	//是否开启debug模式
	if !config.DebugMode {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.New()
	engine.Use(gin.Recovery())

	if config.DebugMode {
		engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		docs.SwaggerInfo.BasePath = authRoutesPrefix
	}

	// session 设置
	store := sessions.NewCookieStore([]byte(config.SessionSecret))
	store.Options(sessions.Options{
		Path:   "/",
		MaxAge: config.SessionExpire, // 设置超时时间为一个小时
	})
	engine.Use(sessions.Sessions(sessionName, store))

	// 日志设置
	engine.Use(middleware.LogMiddleware())

	// 登陆认证相关路由
	authRouterGroup := engine.Group(authRoutesPrefix)
	authRouter(authRouterGroup)
	// 业务路由
	serviceRouterGroup := engine.Group(serviceRoutesPrefix)
	serviceRouterGroup.Use(middleware.SessionAuth())
	userRouter(serviceRouterGroup)

	port := config.ServerPort
	runParams := config.ServerHost + ":" + strconv.Itoa(port)
	log.Println("master server at ", runParams)
	if err := engine.Run(runParams); err != nil {
		log.Error(err)
		return
	}
}
