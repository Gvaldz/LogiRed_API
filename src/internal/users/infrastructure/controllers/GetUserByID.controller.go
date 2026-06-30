package controllers

import (
	"logired/src/internal/users/application"
	"logired/src/internal/users/infrastructure/dto" 
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
// @Security     Bearer
// @Produce      json
// @Param        id path int true "ID del usuario"
// @Success      200 {object} dto.UserResponse "datos del usuario" 
// @Failure      400 {object} map[string]string "ID inválido"
// @Failure      401 {object} map[string]string "no autenticado"
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


	userResponse := dto.UserResponse{
		IdUser:      user.IdUser, 
		Name:        user.Name,
		Lastname:    user.Lastname,
		Email:       user.Email,
		NumberPhone: user.NumberPhone,
		Birthdate:   user.Birthdate,
		UserType:    user.UserType,
		ImageURL:    user.ImageURL,
	}

	c.JSON(http.StatusOK, gin.H{
		"data": userResponse,
	})
}