package services

import (
	"auth-server/dao"
	"auth-server/response"
)

var (
	UserDao dao.UserProxy
	respErr response.ErrorCode
)
