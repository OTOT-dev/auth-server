package api

import (
	"auth-server/model"
	"auth-server/response"
	"github.com/gin-gonic/gin"
)

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
