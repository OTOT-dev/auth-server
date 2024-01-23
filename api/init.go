package api

import (
	"auth-server/response"
	"auth-server/services"
)

var (
	userService services.UserService
	respErr     response.ErrorCode
)
