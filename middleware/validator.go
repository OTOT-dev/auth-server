package middleware

import (
	"auth-server/model"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func Validate[P, Q, B any](c *gin.Context) (params P, query Q, body B) {
	if err := c.ShouldBindUri(&params); err != nil {
		Fail(c, model.ErrParam.AddErr(err))
		return
	}

	if err := c.ShouldBindWith(&query, binding.Query); err != nil {
		Fail(c, model.ErrParam.AddErr(err))
		return
	}

	if c.Request.ContentLength != 0 {
		if err := c.ShouldBindBodyWith(&body, binding.JSON); err != nil {
			Fail(c, model.ErrParam.AddErr(err))
			return
		}
	}

	return
}
