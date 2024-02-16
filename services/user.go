package services

import (
	"auth-server/model"
)

type UserService struct{}

func (UserService) GetUser(userId int64) (user model.User, err model.ErrorCode) {
	user, _, err = proxyUser.GetUserById(userId)
	return
}
