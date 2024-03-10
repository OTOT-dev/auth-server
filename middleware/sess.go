package middleware

import (
	"auth-server/model"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SessionAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")
		if user == nil {
			Fail(c, model.ErrNotLogin)
			return
		}
		c.Next()
	}
}
