package router

import (
	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func userRouter(router *gin.RouterGroup) {
	// user
	router.GET("/users/:id", ginSwagger.WrapHandler(swaggerfiles.Handler), apiUser.GetUser)
	router.PATCH("/users/:id", apiUser.UpdateUser)
	router.DELETE("/users/:id", apiUser.DeleteUser)
}
