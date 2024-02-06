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

func (UserApi) CreateUser(c *gin.Context) {
	// var param model.User
	// if err := c.ShouldBindJSON(&param); err != nil {
	// 	middleware.Fail(c, model.ErrParam.AddErr(err))
	// 	return
	// }

	_, _, param, err := middleware.Validate[any, any, model.User](c)
	if err != nil {
		middleware.Fail(c, model.ErrParam.AddErr(err))
		return
	}

	errCreate := userService.CreateUser(&param)
	middleware.Auto(c, errCreate, nil)
	return
}

func (UserApi) GetUser(c *gin.Context) {
	// id := c.Param("id")
	// if id == "" {
	// 	middleware.Fail(c, model.ErrParam)
	// 	return
	// }
	// userId, err := strconv.ParseInt(id, 10, 64)

	params, _, _, err := middleware.Validate[GetUserParam, any, any](c)
	if err != nil {
		middleware.Fail(c, model.ErrParam.AddErr(err))
		return
	}

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
