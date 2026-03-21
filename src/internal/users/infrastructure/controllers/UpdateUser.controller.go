package controllers

import (
	"fmt"
	"logired/src/internal/users/application"
	"logired/src/internal/users/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateUserController struct {
	updateUserUC *application.UpdateUser
}

func NewUpdateUserController(updateUserUC *application.UpdateUser) *UpdateUserController {
	return &UpdateUserController{
		updateUserUC: updateUserUC,
	}
}

func (c *UpdateUserController) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")

	var user entities.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var idInt int32
	if _, err := fmt.Sscanf(id, "%d", &idInt); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := c.updateUserUC.Execute(idInt, user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Usuario actualizado correctamente"})
}
