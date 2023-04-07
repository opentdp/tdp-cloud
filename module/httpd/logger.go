package httpd

import (
	"time"

	"github.com/gin-gonic/gin"

	"tdp-cloud/helper/logman"
)

// 接收框架默认的日志

func Logger() gin.HandlerFunc {

	logger := logman.Named("gin.access")

	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		c.Next()

		logger.Info(
			path,
			"status", c.Writer.Status(),
			"method", c.Request.Method,
			"path", path,
			"query", query,
			"ip", c.ClientIP(),
			"user-agent", c.Request.UserAgent(),
			"errors", c.Errors.ByType(gin.ErrorTypePrivate).String(),
			"elapsed", time.Since(start),
		)
	}

}
