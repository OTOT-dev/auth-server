package routes

import (
	"auth-server/api"

	"github.com/gin-gonic/gin"
)

var apiUser api.UserApi

func UserRouter(router *gin.RouterGroup) {
	// user
	router.GET("/users/:id", apiUser.GetUser)
	router.POST("/users", apiUser.CreateUser)
	router.PATCH("/users/:id", apiUser.UpdateUser)
	router.DELETE("/users/:id", apiUser.DeleteUser)
}
