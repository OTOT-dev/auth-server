package router

import (
	"auth-server/loader/router/routes"

	"github.com/gin-gonic/gin"
)

func InitRouter(engine *gin.Engine) {
	userRouterGroup := engine.Group("/api/v1")
	routes.UserRouter(userRouterGroup)
}
