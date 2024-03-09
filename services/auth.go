package services

import (
	"auth-server/common"
	"auth-server/model"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
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

func (AuthService) Login(loginUser model.LoginUser, c *gin.Context) (err model.ErrorCode) {
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
	salt := user.Salt
	enPassword := common.MD5(salt + loginUser.Password)
	if enPassword != user.Password {
		err = model.ErrLonginParam
		return
	}
	sess := sessions.Default(c)
	sess.Set("user", user.ID)
	saveErr := sess.Save()
	if saveErr != nil {
		err = model.ErrSession.AddErr(saveErr)
	}
	return
}

func (AuthService) Logout(c *gin.Context) (err model.ErrorCode) {
	sess := sessions.Default(c)
	sess.Clear()
	saveErr := sess.Save()
	if saveErr != nil {
		err = model.ErrSession.AddErr(saveErr)
	}
	return
}
