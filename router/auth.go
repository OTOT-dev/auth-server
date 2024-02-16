package router

import "github.com/gin-gonic/gin"

func authRouter(router *gin.RouterGroup) {
	router.POST("/auth/login", apiAuth.Login)
	router.POST("/auth/register", apiAuth.Register)
}
