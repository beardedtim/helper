package http

import (
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Middleware struct{}

func (middleware *Middleware) SharedHeaders() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Server Information
		ctx.Header("X-Server", "Helper")
		ctx.Header("X-Version", "0.0.1")

		// OWASP Headers Recommendations
		ctx.Header("X-Frame-Options", "deny")
		ctx.Header("X-Content-Type-Options", "nosniff")
		ctx.Header("Content-Security-Policy", "script-src 'self'")
		ctx.Header("X-Permitted-Cross-Domain-Policies", "none")
		ctx.Header("Referrer-Policy", "no-referrer")
		ctx.Header("Cross-Origin-Embedder-Policy", "require-corp")
		ctx.Header("Cross-Origin-Opener-Policy", "same-origin")
		ctx.Header("Cross-Origin-Resource-Policy", "same-origin")
		ctx.Header("X-XSS-Protection", "0")

		ctx.Next()
	}
}

func (middleware *Middleware) Logging() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Starting time request
		startTime := time.Now()

		// Processing request
		ctx.Next()

		// End Time request
		endTime := time.Now()

		// execution time
		latencyTime := endTime.Sub(startTime)

		// Request method
		reqMethod := ctx.Request.Method

		// Request route
		reqUri := ctx.Request.RequestURI

		// status code
		statusCode := ctx.Writer.Status()

		// Request IP
		clientIP := ctx.ClientIP()

		log.WithFields(log.Fields{
			"METHOD":    reqMethod,
			"URI":       reqUri,
			"STATUS":    statusCode,
			"LATENCY":   latencyTime,
			"CLIENT_IP": clientIP,
		}).Info("HTTP REQUEST")
	}
}
