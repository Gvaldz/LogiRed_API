package infrastructure

import (
	"logired/src/core/services/auth/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

type AuthRoutes struct {
	LoginController       *controllers.LoginController
	GoogleLoginController *controllers.GoogleLoginController
}

func NewAuthRoutes(
	loginController *controllers.LoginController,
	googleLoginController *controllers.GoogleLoginController,
) *AuthRoutes {
	return &AuthRoutes{
		LoginController:       loginController,
		GoogleLoginController: googleLoginController,
	}
}

func (r *AuthRoutes) AttachRoutes(router *gin.Engine) {
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/login", r.LoginController.Login)
		authGroup.POST("/google", r.GoogleLoginController.GoogleLogin)
	}
}
