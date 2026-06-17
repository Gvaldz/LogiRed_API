// src/server/middleware/strict_rate_limiter.go
package middleware

import (
	"github.com/gin-gonic/gin"
)

func StrictRateLimit(rps float64, burst int) gin.HandlerFunc {
	rl := NewRateLimiter(rps, burst)
	return rl.Middleware()
}