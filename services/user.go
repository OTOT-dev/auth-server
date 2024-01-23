package services

import (
	"auth-server/common"
	"auth-server/component/response"
	"auth-server/dao"
	"auth-server/model"
)

var userDao dao.UserDao

type UserService struct{}

func (UserService) CreateUser(user *model.UserProps) (err response.ErrorCode) {
	// 产生slat
	slat := common.GetRandomString(16)
	// todo 后续采用加密更强的算法，例如AES
	user.Password = common.MD5(slat)
	createErr := userDao.CreateUser(user)
	if createErr != nil {
		err = response.ErrDbExec
		return
	}
	return
}

func (UserService) GetUser(userId int64) (user *model.UserProps, err response.ErrorCode) {
	var dbErr error
	user, dbErr = userDao.GetUser(userId)

	if dbErr != nil {
		err = response.ErrDbExec
		return
	}
	return
}
