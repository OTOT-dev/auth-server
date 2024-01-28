package services

import (
	"auth-server/common"
	"auth-server/model"
)

type UserService struct{}

func (UserService) CreateUser(user *model.User) (err model.ErrorCode) {
	// 产生slat
	slat := common.GetRandomString(16)
	// todo 后续采用加密更强的算法，例如AES
	user.Password = common.MD5(slat)
	err = proxyUser.CreateUser(user)
	return
}

func (UserService) GetUser(userId int64) (user *model.User, err model.ErrorCode) {
	user, err = proxyUser.GetUser(userId)
	return
}
