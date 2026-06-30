package controllers

import (
    "logired/src/internal/proposals/application"
    "logired/src/internal/proposals/domain/entities"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

// GetProposalsByRideController maneja la obtención de propuestas por viaje
type GetProposalsByRideController struct {
    uc *application.GetProposalsByRide
}

func NewGetProposalsByRideController(uc *application.GetProposalsByRide) *GetProposalsByRideController {
    return &GetProposalsByRideController{uc: uc}
}

// GetByRide godoc
// @Summary      Obtener todas las propuestas de un viaje
// @Tags         proposals
// @Security     Bearer
// @Produce      json
// @Param        idride path int true "ID del viaje"
// @Success      200 {object} map[string]interface{} "lista de propuestas"
// @Failure      400 {object} map[string]string "ID inválido"
// @Failure      401 {object} map[string]string "no autenticado"
// @Failure      403 {object} map[string]string "no eres el cliente del viaje"
// @Failure      500 {object} map[string]string "error interno"
// @Router       /proposals/ride/{idride} [get]
func (ctrl *GetProposalsByRideController) GetByRide(c *gin.Context) {
    idRide, err := strconv.ParseInt(c.Param("idride"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID de viaje inválido"})
        return
    }

    userIDInterface, exists := c.Get("userID")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
        return
    }
    userID := userIDInterface.(int32)

    proposals, err := ctrl.uc.Execute(int32(idRide), userID)
    if err != nil {
        if err.Error() == "solo el cliente dueño del viaje puede ver sus propuestas" {
            c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    if len(proposals) == 0 {
        c.JSON(http.StatusOK, gin.H{
            "message":   "No hay propuestas para este viaje",
            "proposals": []entities.Proposal{},
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "total":     len(proposals),
        "proposals": proposals,
    })
}

type GetProposalsByRideResponse struct {
    Total     int                  `json:"total"`
    Proposals []entities.Proposal  `json:"proposals"`
}