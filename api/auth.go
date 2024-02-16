package api

import (
	"auth-server/middleware"
	"auth-server/model"
	"github.com/gin-gonic/gin"
)

type AuthApi struct{}

func (AuthApi) Login(c *gin.Context) {
	var param model.LoginUser
	if err := c.ShouldBindJSON(&param); err != nil {
		middleware.Fail(c, model.ErrParam.AddErr(err))
		return
	}
	token, err := authService.Login(param)
	middleware.Auto(c, err, map[string]string{
		"token": token,
	})
	return
}

func (AuthApi) Register(c *gin.Context) {
	var param model.User
	if err := c.ShouldBindJSON(&param); err != nil {
		middleware.Fail(c, model.ErrParam.AddErr(err))
		return
	}
	err := authService.Register(&param)
	middleware.Auto(c, err, nil)
	return
}
