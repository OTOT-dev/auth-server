package user

import (
	"auth-server/component/response"
	"auth-server/model"
	"auth-server/services"

	"github.com/gin-gonic/gin"
)

var userService services.UserService

type UserController struct{}

func (UserController) CreateUser(c *gin.Context) {
	var param model.UserProps
	if err := c.ShouldBindJSON(&param); err != nil {
		response.Fail(c, response.ErrParam)
		return
	}
	err := userService.CreateUser(&param)
	response.Auto(c, err, nil)
	return
}

func (UserController) GetUser(gin *gin.Context) {
	println("lalal")
}

func (UserController) UpdateUser(gin *gin.Context) {
}

func (UserController) DeleteUser(gin *gin.Context) {
}
