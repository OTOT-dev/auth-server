package services

import (
	"auth-server/common"
	"auth-server/model"
	"auth-server/response"
)

type UserService struct{}

func (UserService) CreateUser(user *model.User) (err response.ErrorCode) {
	// 产生slat
	slat := common.GetRandomString(16)
	// todo 后续采用加密更强的算法，例如AES
	user.Password = common.MD5(slat)
	createErr := proxyUser.CreateUser(user)
	if createErr != nil {
		err = response.ErrDbExec
		return
	}
	return
}

func (UserService) GetUser(userId int64) (user *model.User, err response.ErrorCode) {
	var dbErr error
	user, dbErr = proxyUser.GetUser(userId)

	if dbErr != nil {
		err = response.ErrDbExec
		return
	}
	return
}
