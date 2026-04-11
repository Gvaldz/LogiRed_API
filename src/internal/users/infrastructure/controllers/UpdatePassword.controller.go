package controllers

import (
	"logired/src/internal/users/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdatePasswordController struct {
	updatePasswordUC *application.UpdatePassword
}

func NewUpdatePasswordController(updatePasswordUC *application.UpdatePassword) *UpdatePasswordController {
	return &UpdatePasswordController{
		updatePasswordUC: updatePasswordUC,
	}
}

func (c *UpdatePasswordController) UpdatePassword(ctx *gin.Context) {
    userIDInterface, exists := ctx.Get("userID")
    if !exists {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
        return
    }
    tokenUserID := userIDInterface.(int32)

    var request struct {

        OldPassword string `json:"oldPassword" binding:"required"`
        NewPassword string `json:"newPassword" binding:"required"`
    }

    if err := ctx.ShouldBindJSON(&request); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
        return
    }

    err := c.updatePasswordUC.Execute(tokenUserID, request.OldPassword, request.NewPassword)
    if err != nil {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "Contraseña actualizada correctamente"})
}