package services

import (
	"auth-server/proxy"
	"auth-server/response"
)

var (
	proxyUser proxy.UserProxy
	respErr   response.ErrorCode
)
