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
// @Description  Devuelve los viajes del usuario autenticado (cliente o conductor) según el estado indicado (1-6)
// @Tags         rides
// @Produce      json
// @Param        idstatus query int true "Estado del viaje (1: asignado, 2: en camino, 3: recogido, 4: completado, 5: cancelado, 6: disponible)"
// @Security     ApiKeyAuth
// @Success      200 {object} map[string]interface{} "lista de viajes"
// @Failure      400 {object} map[string]string "idstatus requerido o inválido"
// @Failure      401 {object} map[string]string "no autenticado"
// @Failure      500 {object} map[string]string "error interno"
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