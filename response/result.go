package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code,omitempty"` // 错误码
	Msg  string      `json:"msg,omitempty"`  // 错误描述
	Data interface{} `json:"data,omitempty"` // 返回数据
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		0,
		"",
		data,
	})
}

func Auto(c *gin.Context, err ErrorCode, data interface{}) {
	c.JSON(http.StatusOK, Response{
		err.Code,
		err.Msg,
		data,
	})
}

func Fail(c *gin.Context, error ErrorCode) {
	c.JSON(http.StatusOK, Response{
		error.Code,
		error.Msg,
		nil,
	})
}
