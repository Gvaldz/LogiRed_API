package controllers

import (
    "logired/src/internal/rides/application"
    userEntities "logired/src/internal/users/domain/entities"
    "logired/src/internal/rides/domain/entities"
    "net/http"
    "github.com/gin-gonic/gin"
)

type GetRidesHistoryController struct {
    getRidesHistory *application.GetRidesHistory
}

func NewGetRidesHistoryController(uc *application.GetRidesHistory) *GetRidesHistoryController {
    return &GetRidesHistoryController{getRidesHistory: uc}
}

// GetHistory godoc
// @Summary      Obtener historial de viajes (completados o cancelados)
// @Description  Devuelve viajes con estado 4 o 5 (historial) para el usuario autenticado (cliente o conductor)
// @Tags         rides
// @Produce      json
// @Security     ApiKeyAuth
// @Success      200 {object} map[string]interface{} "lista de viajes históricos"
// @Failure      401 {object} map[string]string "no autenticado"
// @Failure      500 {object} map[string]string "error interno"
// @Router       /rides/history [get]
func (ctrl *GetRidesHistoryController) GetHistory(c *gin.Context) {
    userIDInterface, exists := c.Get("userID")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
        return
    }
    userID := userIDInterface.(int32)

    userInterface, exists := c.Get("user")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
        return
    }
    user := userInterface.(userEntities.User)

    rides, err := ctrl.getRidesHistory.Execute(userID, int32(user.UserType))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    if len(rides) == 0 {
        c.JSON(http.StatusOK, gin.H{
            "message": "No hay viajes en el historial",
            "rides":   []entities.Ride{},
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "total": len(rides),
        "rides": rides,
    })
}