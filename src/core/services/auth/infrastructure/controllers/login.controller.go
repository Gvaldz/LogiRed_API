package controllers

import (
	"logired/src/core/services/auth/application"
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

// Login godoc
// @Summary      Iniciar sesión
// @Description  Autentica a un usuario y devuelve un token JWT
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        body body entities.User true "Credenciales" Example({"email":"correo@example.com","password":"123456"})
// @Success      200 {object} map[string]interface{} "token y expiración"
// @Failure      400 {object} map[string]string "petición inválida"
// @Failure      401 {object} map[string]string "credenciales incorrectas"
// @Router       /auth/login [post]
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
