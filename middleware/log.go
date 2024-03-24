package middleware

import (
	"auth-server/common"
	"auth-server/config"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"time"
)

var log = common.Log

const (
	statusColor200 = 42
	statusColor404 = 43
	statusColor500 = 41
	statusColor400 = 40

	methodColorGET   = 44
	methodColorPOST  = 45
	methodColorPATCH = 46
	methodColorDELTE = 47
)

func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		params := ""
		if config.DebugMode && c.ContentType() == "application/json" {
			ByteBody, _ := io.ReadAll(c.Request.Body)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(ByteBody))
			params = string(ByteBody)
		}

		// Process request
		c.Next()

		// Stop timer
		end := time.Now()
		timeSub := end.Sub(start).Seconds()
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()

		if raw != "" {
			path = path + "?" + raw
		}

		var statusColor string
		switch statusCode {
		case http.StatusOK:
			statusColor = fmt.Sprintf("\033[%dm %d \033[0m", statusColor200, statusCode)
		case http.StatusNotFound:
			statusColor = fmt.Sprintf("\033[%dm %d \033[0m", statusColor404, statusCode)
		case http.StatusBadRequest:
			statusColor = fmt.Sprintf("\033[%dm %d \033[0m", statusColor400, statusCode)
		case http.StatusInternalServerError:
			statusColor = fmt.Sprintf("\033[%dm %d \033[0m", statusColor500, statusCode)
		}

		var methodColor string
		switch method {
		case http.MethodGet:
			methodColor = fmt.Sprintf("\033[%dm %s \033[0m", methodColorGET, method)
		case http.MethodDelete:
			methodColor = fmt.Sprintf("\033[%dm %s \033[0m", methodColorDELTE, method)
		case http.MethodPost:
			methodColor = fmt.Sprintf("\033[%dm %s \033[0m", methodColorPOST, method)
		case http.MethodPatch:
			methodColor = fmt.Sprintf("\033[%dm %s \033[0m", methodColorPATCH, method)
		}
		// 写入到文件中不包含相关到颜色编码，所以需要额外到字段来存储原始信息
		fields := logrus.Fields{
			"time":     start.Format("2006-01-02 15:04:06"),
			"method":   method,
			"status":   statusCode,
			"path":     path,
			"clientIp": clientIP,
			"timeSub":  timeSub,
		}
		if config.DebugMode {
			log.WithFields(fields).Infof("[GIN] %s |%s| %s| %s | %s | %fs \n %s",
				start.Format("2006-01-02 15:04:06"),
				statusColor,
				clientIP,
				methodColor,
				path,
				timeSub,
				params,
			)
		} else {
			log.WithFields(fields).Infof("[GIN] %s |%s| %s| %s | %s | %fs",
				start.Format("2006-01-02 15:04:06"),
				statusColor,
				clientIP,
				methodColor,
				path,
				timeSub,
			)
		}

	}
}
