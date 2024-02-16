package services

import (
	"auth-server/common"
	"auth-server/config"
	"auth-server/model"
	"github.com/pkg/errors"
)

type AuthService struct{}

var (
	UserNotFoundErr = errors.New("user not found")
	UserExistErr    = errors.New("user exist")
)

func (AuthService) Register(user *model.User) (err model.ErrorCode) {
	var found bool
	_, found, err = proxyUser.GetUserByUsername(user.Username)
	if err.Code != 0 {
		return
	}
	if found {
		err = model.ErrRegisterParam.AddErr(UserExistErr)
		return
	}

	// 产生slat
	slat := common.GetRandomString(16)
	user.Salt = slat
	// todo 后续采用加密更强的算法，例如AES
	user.Password = common.MD5(slat + user.Password)
	err = proxyUser.CreateUser(user)
	return
}

func (AuthService) Login(loginUser model.LoginUser) (token string, err model.ErrorCode) {
	var user model.User
	var found bool
	user, found, err = proxyUser.GetUserByUsername(loginUser.Username)
	if err.Code != 0 {
		return
	}
	if !found {
		err = model.ErrLonginParam.AddErr(UserNotFoundErr)
		return
	}
	slat := user.Salt
	enPassword := common.MD5(slat + loginUser.Password)
	if enPassword != user.Password {
		err = model.ErrLonginParam
		return
	}
	var genTokenErr error
	token, genTokenErr = common.GenerateToken(loginUser.Username, config.JwtSecret, config.JwtExpire, config.ServerName)
	if genTokenErr != nil {
		err = model.ErrGenToken.AddErr(genTokenErr)
		return
	}
	return
}
