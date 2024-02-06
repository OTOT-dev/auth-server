package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func Validate[P, Q, B any](c *gin.Context) (params P, query Q, body B, err error) {
	if err = c.ShouldBindUri(&params); err != nil {
		return
	}

	if err = c.ShouldBindWith(&query, binding.Query); err != nil {
		return
	}

	if c.Request.ContentLength != 0 {
		if err = c.ShouldBindBodyWith(&body, binding.JSON); err != nil {
			return
		}
	}

	return
}
