package middleware

import (
	"auth-server/model"
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ValidatorMiddleware(pType, qType, bType reflect.Type) gin.HandlerFunc {
	return func(c *gin.Context) {
		vd := validator.New()

		if pType != nil {
			params := reflect.New(pType).Interface()
			if err := c.ShouldBindUri(params); err != nil {
				Fail(c, model.ErrParam.AddErr(err))
				return
			}
			if err := vd.Struct(params); err != nil {
				Fail(c, model.ErrParam.AddErr(err))
				return
			}
		}

		if qType != nil {
			query := reflect.New(qType).Interface()
			if err := c.ShouldBindQuery(query); err != nil {
				Fail(c, model.ErrParam.AddErr(err))
			}
			fmt.Println(query)
			if err := vd.Struct(query); err != nil {
				Fail(c, model.ErrParam.AddErr(err))
				return
			}
		}

		if bType != nil {
			body := reflect.New(bType).Interface()
			if err := c.ShouldBindJSON(body); err != nil {
				Fail(c, model.ErrParam.AddErr(err))
				return
			}
			if err := vd.Struct(body); err != nil {
				Fail(c, model.ErrParam.AddErr(err))
				return
			}
		}
	}
}
