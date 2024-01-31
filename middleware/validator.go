package middleware

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ValidatorMiddleware(pType, qType, bType reflect.Type) gin.HandlerFunc {
	return func(c *gin.Context) {
		vd := validator.New()

		if bType != nil {
			body := reflect.New(bType)
			println("111")
			println(body.String)
			c.ShouldBindJSON(body)
			err := vd.Struct(body)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			if err != nil {
				println(err.Error())
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
		}
	}
}
