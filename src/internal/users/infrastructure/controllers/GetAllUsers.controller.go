package controllers

import (
	"logired/src/internal/users/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllUsersController struct {
	GetAllUsers *application.GetAllUsers
}

func NewGetAllUsersController(getAllUsers *application.GetAllUsers) *GetAllUsersController {
	return &GetAllUsersController{
		GetAllUsers: getAllUsers,
	}
}
// GetAll godoc
// @Summary      Obtener todos los usuarios (solo admin)
// @Description  Devuelve lista de usuarios (solo para administradores)
// @Tags         users
// @Produce      json
// @Security     Bearer
// @Success      200 {array} entities.User "lista de usuarios"
// @Failure      401 {object} map[string]string "no autenticado"
// @Failure      403 {object} map[string]string "no autorizado"
// @Router       /users [get]
func (h *GetAllUsersController) GetAll(c *gin.Context) {
	doctors, err := h.GetAllUsers.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, doctors)
}
