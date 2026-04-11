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