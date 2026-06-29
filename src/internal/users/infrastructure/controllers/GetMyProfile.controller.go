package controllers

import (
    "logired/src/internal/users/application"
    "net/http"

    "github.com/gin-gonic/gin"
)

type GetMyProfileController struct {
    getMyProfile *application.GetMyProfile
}

func NewGetMyProfileController(uc *application.GetMyProfile) *GetMyProfileController {
    return &GetMyProfileController{getMyProfile: uc}
}

// GetMyProfile godoc
// @Summary      Obtener perfil del usuario autenticado
// @Description  Devuelve todos los datos del usuario logueado. Si es conductor incluye calificación y vehículo. Incluye total de viajes realizados.
// @Tags         users
// @Security     Bearer
// @Produce      json
// @Success      200 {object} entities.MyProfile "perfil propio"
// @Failure      401 {object} map[string]string "no autenticado"
// @Failure      404 {object} map[string]string "usuario no encontrado"
// @Router       /users/me [get]
func (ctrl *GetMyProfileController) GetMyProfile(c *gin.Context) {
    userIDInterface, exists := c.Get("userID")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
        return
    }
    userID := userIDInterface.(int32)

    profile, err := ctrl.getMyProfile.Execute(userID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, profile)
}
