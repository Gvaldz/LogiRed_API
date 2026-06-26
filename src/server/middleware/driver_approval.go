package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	userEntities "logired/src/internal/users/domain/entities"
)

const UserTypeConductor = 2

func RequireApprovedDriver(authRepo interface{}) gin.HandlerFunc {
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

		// El approved ya fue extraído del token por AuthMiddleware
		approved, exists := c.Get("approved")
		if !exists {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "error al verificar aprobación"})
			return
		}

		if !approved.(bool) {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error":    "tu cuenta de conductor está pendiente de aprobación",
				"approved": false,
			})
			return
		}

		c.Set("approved", true)
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
