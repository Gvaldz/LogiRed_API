package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	tokenService "logired/src/internal/services/auth/domain"
	users_domain "logired/src/internal/users/domain"
)

func AuthMiddleware(tokenService tokenService.TokenService, userRepo users_domain.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token de autorización requerido"})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "formato de token inválido"})
			return
		}

		userID, err := tokenService.ValidateToken(parts[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
			return
		}

		user, err := userRepo.GetUserByID(userID)
		
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "error al verificar usuario"})
			return
		}

		c.Set("userID", userID)
		c.Set("user", user)
		c.Next()
	}
}
