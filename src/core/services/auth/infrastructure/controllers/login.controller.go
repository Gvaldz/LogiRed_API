package controllers

import (
	"logired/src/core/services/auth/application"
	"logired/src/internal/users/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

// LoginRequest es la estructura que Swagger mostrará
type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginController struct {
	loginUC *application.Login
}

func NewLoginController(loginUC *application.Login) *LoginController {
	return &LoginController{loginUC: loginUC}
}

// Login godoc
// @Summary      Iniciar sesión
// @Description  Autentica a un usuario y devuelve un token JWT que contiene usertype y approved (para conductores)
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        body body LoginRequest true "Credenciales" Example({"email":"correo@example.com","password":"123456"})
// @Success      200 {object} map[string]interface{} "expires_at: timestamp del token"
// @Failure      400 {object} map[string]string "petición inválida"
// @Failure      401 {object} map[string]string "credenciales incorrectas"
// @Router       /auth/login [post]
func (c *LoginController) Login(ctx *gin.Context) {
	// Usamos LoginRequest en lugar de entities.User
	var credentials LoginRequest
	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "petición inválida"})
		return
	}

	// Construimos el objeto User que espera el caso de uso
	user := entities.User{
		Email:    credentials.Email,
		Password: credentials.Password,
	}

	result, err := c.loginUC.Execute(user)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// El token contiene usertype y approved (para conductores)
	// NO se devuelven en el body, se extraen del token
	ctx.Header("Authorization", "Bearer "+result.Token.Token)
	ctx.JSON(http.StatusOK, gin.H{
		"token":      result.Token.Token,
		"expires_at": result.Token.ExpiresAt,
	})
}
