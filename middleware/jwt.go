package middleware

import (
	"auth-server/common"
	"auth-server/config"
	"auth-server/model"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err model.ErrorCode
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			err = model.ErrParam.AddErr(errors.New("Authorization为空"))
		} else {
			_, parseErr := common.ParseToken(token, config.JwtSecret)
			if parseErr != nil {
				err = model.ErrAuthToken.AddErr(parseErr)
			}
		}
		if err.Code != 0 {
			Fail(c, err)
			return
		}

		c.Next()
	}
}
