package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

func NonBlockingRateLimitMiddleware(rateLimit *ratelimit.Bucket) gin.HandlerFunc {
	return func(c *gin.Context) {
		// If there is no token available, deny the request immediately.
		if rateLimit.TakeAvailable(1) == 0 {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "Too many requests. Please try again later.",
			})
			return
		}

		c.Next()
	}
}
