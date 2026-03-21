package controllers

import (
	"logired/src/internal/services/auth/application"
	"logired/src/internal/users/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	loginUC *application.Login
}

func NewLoginController(loginUC *application.Login) *LoginController {
	return &LoginController{loginUC: loginUC}
}

func (c *LoginController) Login(ctx *gin.Context) {
	var credentials entities.User
	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "petición inválida"})
		return
	}

	token, err := c.loginUC.Execute(credentials)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.Header("Authorization", "Bearer "+token.Token)
	ctx.JSON(http.StatusOK, gin.H{
		"expires_at": token.ExpiresAt,
	})
}
