package midware

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

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
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("elapsed", time.Since(start)),
		)
	}

}
