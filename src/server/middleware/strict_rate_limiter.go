package middleware

import (
	"github.com/gin-gonic/gin"
)

func StrictRateLimit(rps float64, burst int) gin.HandlerFunc {
	rl := NewRateLimiter(rps, burst)
	return rl.Middleware()
}