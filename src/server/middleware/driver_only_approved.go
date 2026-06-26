package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	authDomain "logired/src/core/services/auth/domain"
	userEntities "logired/src/internal/users/domain/entities"
)

// RequireOnlyApprovedDriver - Middleware que RECHAZA a clientes y conductores no aprobados
// Solo permite CONDUCTORES APROBADOS
func RequireOnlyApprovedDriver(authRepo authDomain.AuthRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		userInterface, exists := c.Get("user")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "usuario no autenticado"})
			return
		}

		user, ok := userInterface.(userEntities.User)
		if !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "error al procesar usuario"})
			return
		}

		// Si NO es conductor, rechaza
		if user.UserType != UserTypeConductor {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error":         "solo los conductores pueden acceder a este recurso",
				"required_role": "driver",
			})
			return
		}

		// Si es conductor, verifica aprobación
		approved, err := authRepo.FindDriverApproved(user.IdUser)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "error al verificar estado de aprobación"})
			return
		}

		if !approved {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error":    "tu cuenta de conductor está pendiente de aprobación. Por favor, espera a que sea aprobada",
				"approved": false,
			})
			return
		}

		c.Set("approved", true)
		c.Next()
	}
}
