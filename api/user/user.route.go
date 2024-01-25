package user

import (
	"github.com/gin-gonic/gin"
)

var userController UserController

func RegisterUserRouter(rg *gin.RouterGroup) {
	// user
	rg.GET("/users/:id", userController.GetUser)
	rg.POST("/users", userController.CreateUser)
	rg.PATCH("/users/:id", userController.UpdateUser)
	rg.DELETE("/users/:id", userController.DeleteUser)
}
