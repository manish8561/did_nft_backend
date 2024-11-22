package middlewares

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// RateLimitMiddleware implements rate limiting using token bucket algorithm
func RateLimitMiddleware(maxRequests int, timeWindow time.Duration) gin.HandlerFunc {
	tokens := maxRequests
	lastRefillTime := time.Now()

	return func(c *gin.Context) {
		// Refill the bucket if the time window has passed
		now := time.Now()
		if now.Sub(lastRefillTime) > timeWindow {
			tokens = maxRequests
			lastRefillTime = now
		}

		// Check if there are tokens available
		if tokens > 0 {
			tokens--
			c.Next()
		} else {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "Too many requests. Please try again later.",
			})
			c.Abort()
		}
	}
}
