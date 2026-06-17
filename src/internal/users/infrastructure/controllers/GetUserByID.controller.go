package controllers

import (
	"logired/src/internal/users/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetByUserIDController struct {
	getByUserID *application.GetUserByID
}

func NewGetByUserIDController(getByUserID *application.GetUserByID) *GetByUserIDController {
	return &GetByUserIDController{getByUserID: getByUserID}
}

// GetByUserID godoc
// @Summary      Obtener usuario por ID (versión simplificada)
// @Tags         users
// @Produce      json
// @Param        id path int true "ID del usuario"
// @Success      200 {object} map[string]interface{} "datos del usuario"
// @Failure      400 {object} map[string]string "ID inválido"
// @Failure      404 {object} map[string]string "usuario no encontrado"
// @Router       /users/{id} [get]
func (h *GetByUserIDController) GetByUserID(c *gin.Context) {
	iduser := c.Param("id")
	idInt, err := strconv.Atoi(iduser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario inválido"})
		return
	}

	user, err := h.getByUserID.Execute(int32(idInt))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
