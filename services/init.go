package services

import (
	"auth-server/component/response"
	"auth-server/dao"
)

var (
	UserDao dao.UserProxy
	respErr response.ErrorCode
)
