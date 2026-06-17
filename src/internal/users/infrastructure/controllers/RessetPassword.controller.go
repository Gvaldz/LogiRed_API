package controllers

import (
	"logired/src/internal/users/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RessetPasswordController struct {
	ressetPasswordUC *application.RessetPassword
}

func NewRessetPasswordController(ressetPasswordUC *application.RessetPassword) *RessetPasswordController {
	return &RessetPasswordController{
		ressetPasswordUC: ressetPasswordUC,
	}
}
// RessetPassword godoc
// @Summary      Restablecer contraseña (por email)
// @Description  Actualiza la contraseña de un usuario dado su email
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        body body object true "Datos" Example({"email":"correo@example.com","newPassword":"nueva123"})
// @Success      200 {object} map[string]string "mensaje de éxito"
// @Failure      400 {object} map[string]string "datos inválidos"
// @Failure      401 {object} map[string]string "error"
// @Router       /users/password-reset [put]
func (c *RessetPasswordController) RessetPassword(ctx *gin.Context) {
    var request struct {
        Email       string `json:"email" binding:"required,email"`
        NewPassword string `json:"newPassword" binding:"required"`
    }

    if err := ctx.ShouldBindJSON(&request); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
        return
    }

    err := c.ressetPasswordUC.Execute(request.Email, request.NewPassword)
    if err != nil {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "Contraseña actualizada correctamente"})
}