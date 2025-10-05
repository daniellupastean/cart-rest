package internal

import (
	"time"

	"github.com/gin-gonic/gin"
)

func MetricsMiddleware(telemetry *Telemetry) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/metrics" {
			c.Next()
			return
		}

		start := time.Now()
		c.Next()

		duration := time.Since(start)
		method := c.Request.Method
		path := c.FullPath()
		if path == "" {
			path = c.Request.URL.Path
		}
		statusCode := c.Writer.Status()

		telemetry.RecordRequest(method, path, statusCode, duration)
	}
}

