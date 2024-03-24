package middleware

import (
	"auth-server/common"
	"auth-server/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type response struct {
	Code int         `json:"code,omitempty"` // 错误码
	Msg  string      `json:"msg,omitempty"`  // 错误信息
	Data interface{} `json:"data,omitempty"` // 返回数据
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, response{
		0,
		"",
		data,
	})
}

// Auto 根据传入的内容自动推断是属于正确还是错误
func Auto(c *gin.Context, err model.ErrorCode, data interface{}) {
	var resp response
	if err.Code != 0 {
		resp.Code = err.Code
		resp.Msg = err.Msg
		if err.Err != nil {
			resp.Data = err.Err.Error()
		}
		common.Log.Error(resp)
	} else {
		resp.Data = data
	}
	c.AbortWithStatusJSON(http.StatusOK, resp)
}

func Fail(c *gin.Context, error model.ErrorCode) {
	var errorData string
	//当未传入自定义错误时为空
	if error.Err != nil {
		errorData = error.Err.Error()
	}
	resp := response{
		Code: error.Code,
		Msg:  error.Msg,
		Data: errorData,
	}
	common.Log.Error(resp)
	c.AbortWithStatusJSON(http.StatusOK, resp)
}
