package controllers

import (
    "logired/src/internal/proposals/application"
    "logired/src/internal/proposals/domain/entities"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type GetProposalsByRideController struct {
    uc *application.GetProposalsByRide
}

func NewGetProposalsByRideController(uc *application.GetProposalsByRide) *GetProposalsByRideController {
    return &GetProposalsByRideController{uc: uc}
}

func (ctrl *GetProposalsByRideController) GetByRide(c *gin.Context) {
    idRide, err := strconv.ParseInt(c.Param("idride"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID de viaje inválido"})
        return
    }

    proposals, err := ctrl.uc.Execute(int32(idRide))
    if err != nil {
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