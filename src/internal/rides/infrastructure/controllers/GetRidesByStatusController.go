package controllers

import (
    "logired/src/internal/rides/application"
    userEntities "logired/src/internal/users/domain/entities"
    "logired/src/internal/rides/domain/entities"
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
)

type GetRidesByStatusController struct {
    getRidesByStatus *application.GetRidesByStatus
}

func NewGetRidesByStatusController(uc *application.GetRidesByStatus) *GetRidesByStatusController {
    return &GetRidesByStatusController{getRidesByStatus: uc}
}

// GetByStatus godoc
// @Summary      Obtener viajes por estado
// @Tags         rides
// @Produce      json
// @Param        idstatus query int true "Estado (1-6)"
// @Security     ApiKeyAuth
// @Success      200 {object} map[string]interface{} "total, rides"
// @Failure      400 {object} map[string]string
// @Failure      401 {object} map[string]string
// @Router       /rides/status [get]
func (ctrl *GetRidesByStatusController) GetByStatus(c *gin.Context) {
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

    idStatusStr := c.Query("idstatus")
    if idStatusStr == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "idstatus es requerido"})
        return
    }
    idStatus, err := strconv.ParseInt(idStatusStr, 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "idstatus inválido"})
        return
    }

    if idStatus < 1 || idStatus > 6 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "idstatus debe ser entre 1 y 6"})
        return
    }

    rides, err := ctrl.getRidesByStatus.Execute(userID, int32(user.UserType), int32(idStatus))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    if len(rides) == 0 {
        c.JSON(http.StatusOK, gin.H{
            "message": "No hay viajes con ese estado",
            "rides":   []entities.Ride{},
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "total": len(rides),
        "rides": rides,
    })
}