package user

import (
	"auth-server/component/response"
	"auth-server/model"
	"auth-server/services"
	"strconv"

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

func (UserController) GetUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.Fail(c, response.ErrParam)
		return
	}
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.Fail(c, response.ErrParam)
		return
	}
	user, err1 := userService.GetUser(userId)
	response.Auto(c, err1, user)
	return
}

func (UserController) UpdateUser(gin *gin.Context) {
}

func (UserController) DeleteUser(gin *gin.Context) {
}
