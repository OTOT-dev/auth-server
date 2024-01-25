package loader

import (
	ApiUser "auth-server/api/user"

	"github.com/gin-gonic/gin"
)

func InitRouter(engine *gin.Engine) {
	rg := engine.Group("/api/v1")
	ApiUser.RegisterUserRouter(rg)
}
