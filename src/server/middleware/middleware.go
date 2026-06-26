package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	tokenService "logired/src/core/services/auth/domain"
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

		tokenString := parts[1]

		userID, err := tokenService.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
			return
		}

		user, err := userRepo.GetUserByID(userID)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "error al verificar: " + err.Error()})
			return
		}

		// Extrae usertype del token
		userType, err := tokenService.ExtractUserType(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "error al extraer tipo de usuario del token"})
			return
		}

		// Extrae approved del token
		approved, err := tokenService.ExtractApproved(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "error al extraer estado de aprobación del token"})
			return
		}

		// Almacena en el contexto
		c.Set("userID", userID)
		c.Set("user", user)
		c.Set("userType", userType)
		c.Set("approved", approved)
		c.Next()
	}
}
