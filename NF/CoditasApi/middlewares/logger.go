package middlewares

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// LatencyLoggerMiddleware logs the api response time
func LatencyLoggerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next()
		latency := time.Since(start)
		log.Printf("Request to %s took %v", ctx.Request.URL.Path, latency)
	}
}
