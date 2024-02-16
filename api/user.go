package api

import (
	"auth-server/middleware"
	"auth-server/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetUserParam struct {
	Id string `uri:"id" binding:"required"`
}

type UserApi struct{}

func (UserApi) GetUser(c *gin.Context) {
	params, _, _ := middleware.Validate[GetUserParam, any, any](c)

	userId, err := strconv.ParseInt(params.Id, 10, 64)
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
