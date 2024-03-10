package router

import "github.com/gin-gonic/gin"

func authRouter(router *gin.RouterGroup) {
	router.POST("/login", apiAuth.Login)
	router.POST("/register", apiAuth.Register)
	router.POST("/logout", apiAuth.Logout)
}
