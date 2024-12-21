package api

import (
	"github.com/axidex/CryptBot/pkg/logger"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
	"time"
)

var limiter = rate.NewLimiter(1, 5)

func rateLimiter(c *gin.Context) {
	if !limiter.Allow() {
		c.JSON(http.StatusTooManyRequests, gin.H{"error": "too many requests"})
		c.Abort()
		return
	}
	c.Next()
}

func LoggerMiddleware(logger logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Stop timer
		end := time.Now()
		latency := end.Sub(start)

		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()

		if raw != "" {
			path = path + "?" + raw
		}

		if len(c.Errors.Errors()) > 0 {
			logger.Errorf(
				"path: %s | statusCode: %d | method: %s | clientIP: %s | latency: %s | errorMessage: %s",
				path,
				statusCode,
				method,
				clientIP,
				latency.String(),
				c.Errors.String(),
			)
		}

		logger.Infof(
			"path: %s | statusCode: %d | method: %s | clientIP: %s | latency: %s",
			path,
			statusCode,
			method,
			clientIP,
			latency.String(),
		)

	}
}

// CustomRecoveryFunc is a custom recovery function
func CustomRecoveryFunc(logger logger.Logger) gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		logger.Fatalf(
			"Panic recovered | %v | path: %s | method: %s | ip: %s",
			recovered,
			c.Request.URL.Path,
			c.Request.Method,
			c.ClientIP(),
		)

		// Abort the request and return a 500 Internal Server Error
		c.AbortWithStatus(http.StatusInternalServerError)
	})
}
