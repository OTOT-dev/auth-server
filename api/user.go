package api

import (
	"auth-server/component/response"
	"auth-server/model"
	"auth-server/services"

	"github.com/gin-gonic/gin"
)

var userService services.UserService

type UserApi struct{}

func (UserApi) CreateUser(c *gin.Context) {
	var param model.User
	if err := c.ShouldBindJSON(&param); err != nil {
		response.Fail(c, response.ErrParam)
		return
	}
	err := userService.CreateUser(&param)
	response.Auto(c, err, nil)
	return
}

func (UserApi) GetUser(gin *gin.Context) {
}

func (UserApi) UpdateUser(gin *gin.Context) {
}

func (UserApi) DeleteUser(gin *gin.Context) {
}
