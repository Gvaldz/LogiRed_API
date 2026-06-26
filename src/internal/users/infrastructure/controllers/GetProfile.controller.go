package controllers

import (
    "fmt"
    "logired/src/internal/users/application"
    "net/http"

    "github.com/gin-gonic/gin"
)

type GetUserProfileController struct {
    getProfile *application.GetUserProfile
}

func NewGetUserProfileController(uc *application.GetUserProfile) *GetUserProfileController {
    return &GetUserProfileController{getProfile: uc}
}
// GetProfile godoc
// @Summary      Obtener perfil completo de un usuario (con datos de conductor si aplica)
// @Tags         users
// @Security     Bearer
// @Produce      json
// @Param        id path int true "ID del usuario"
// @Success      200 {object} entities.UserProfile "perfil completo"
// @Failure      400 {object} map[string]string "ID inválido"
// @Failure      401 {object} map[string]string "no autenticado"
// @Failure      404 {object} map[string]string "usuario no encontrado"
// @Router       /users/profile/{id} [get]
func (ctrl *GetUserProfileController) GetProfile(c *gin.Context) {
    var idInt int32
    if _, err := fmt.Sscanf(c.Param("id"), "%d", &idInt); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
    }

    profile, err := ctrl.getProfile.Execute(idInt)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, profile)
}