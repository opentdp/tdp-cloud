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
			logman.Int("status", c.Writer.Status()),
			logman.String("method", c.Request.Method),
			logman.String("path", path),
			logman.String("query", query),
			logman.String("ip", c.ClientIP()),
			logman.String("user-agent", c.Request.UserAgent()),
			logman.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			logman.Duration("elapsed", time.Since(start)),
		)
	}

}
