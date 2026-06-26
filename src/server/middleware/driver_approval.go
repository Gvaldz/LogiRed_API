package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	authDomain "logired/src/core/services/auth/domain"
	userEntities "logired/src/internal/users/domain/entities"
)

const UserTypeConductor = 2

func RequireApprovedDriver(authRepo authDomain.AuthRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		userInterface, exists := c.Get("user")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "usuario no autenticado"})
			return
		}

		user, ok := userInterface.(userEntities.User)
		if !ok || user.UserType != UserTypeConductor {
			c.Next()
			return
		}

		approved, err := authRepo.FindDriverApproved(user.IdUser)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "error al verificar aprobacion"})
			return
		}

		if !approved {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error":    "tu cuenta de conductor esta pendiente de aprobacion",
				"approved": false,
			})
			return
		}

		c.Next()
	}
}

func Chain(handlers ...gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, h := range handlers {
			h(c)
			if c.IsAborted() {
				return
			}
		}
	}
}
