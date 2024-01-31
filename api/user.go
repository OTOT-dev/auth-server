package api

import (
	"auth-server/middleware"
	"auth-server/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserApi struct{}

func (UserApi) CreateUser(c *gin.Context) {
	var param model.User
	if err := c.ShouldBindJSON(&param); err != nil {
		middleware.Fail(c, model.ErrParam.AddErr(err))
		return
	}
	err := userService.CreateUser(&param)
	middleware.Auto(c, err, nil)
	return
}

func (UserApi) GetUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		middleware.Fail(c, model.ErrParam)
		return
	}
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		middleware.Fail(c, model.ErrParam.AddErr(err))
		return
	}
	user, err1 := userService.GetUser(userId)
	middleware.Auto(c, err1, user)
	return
}

func (UserApi) UpdateUser(gin *gin.Context) {
}

func (UserApi) DeleteUser(gin *gin.Context) {
}
