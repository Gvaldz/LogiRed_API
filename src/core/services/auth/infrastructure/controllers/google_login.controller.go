package controllers

import (
	"logired/src/core/services/auth/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GoogleLoginController struct {
	googleLoginUC *application.GoogleLogin
}

func NewGoogleLoginController(uc *application.GoogleLogin) *GoogleLoginController {
	return &GoogleLoginController{googleLoginUC: uc}
}

// GoogleLogin godoc
// @Summary      Iniciar sesión con Google (Firebase)
// @Description  Recibe un Firebase ID Token, lo valida y devuelve un JWT igual que /auth/login
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        body body object true "Firebase ID Token" Example({"idToken":"eyJhbG..."})
// @Success      200 {object} map[string]interface{} "expires_at del token"
// @Failure      400 {object} map[string]string "petición inválida"
// @Failure      401 {object} map[string]string "token de Google inválido"
// @Failure      404 {object} map[string]string "usuario no registrado"
// @Router       /auth/google [post]
func (ctrl *GoogleLoginController) GoogleLogin(c *gin.Context) {
	var body struct {
		IdToken string `json:"idToken" binding:"required"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "idToken es requerido"})
		return
	}

	token, err := ctrl.googleLoginUC.Execute(c.Request.Context(), body.IdToken)
	if err != nil {
		errMsg := err.Error()
		if errMsg == "no existe una cuenta registrada con este correo" {
			c.JSON(http.StatusNotFound, gin.H{"error": errMsg})
			return
		}
		c.JSON(http.StatusUnauthorized, gin.H{"error": errMsg})
		return
	}

	c.Header("Authorization", "Bearer "+token.Token)
	c.JSON(http.StatusOK, gin.H{
		"expires_at": token.ExpiresAt,
	})
}
