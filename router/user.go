package router

import (
	"github.com/gin-gonic/gin"
)

func userRouter(router *gin.RouterGroup) {
	// user
	router.GET("/users/:id", apiUser.GetUser)
	router.PATCH("/users/:id", apiUser.UpdateUser)
	router.DELETE("/users/:id", apiUser.DeleteUser)
}
