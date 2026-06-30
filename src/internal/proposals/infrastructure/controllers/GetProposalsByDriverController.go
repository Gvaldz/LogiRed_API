package controllers

import (
	"logired/src/internal/proposals/application"
	"logired/src/internal/proposals/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetProposalsByDriverController struct {
	uc *application.GetProposalsByDriver
}

func NewGetProposalsByDriverController(uc *application.GetProposalsByDriver) *GetProposalsByDriverController {
	return &GetProposalsByDriverController{uc: uc}
}

// GetMyProposals godoc
// @Summary      Obtener propuestas del conductor autenticado
// @Description  Devuelve todas las propuestas enviadas por el conductor logueado.
// @Tags         proposals
// @Security     Bearer
// @Produce      json
// @Success      200 {object} map[string]interface{} "lista de propuestas"
// @Failure      401 {object} map[string]string "no autenticado"
// @Failure      500 {object} map[string]string "error interno"
// @Router       /proposals/driver [get]
func (ctrl *GetProposalsByDriverController) GetMyProposals(c *gin.Context) {
	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}
	userID := userIDInterface.(int32)

	proposals, err := ctrl.uc.Execute(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(proposals) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message":   "No tienes propuestas registradas",
			"proposals": []entities.Proposal{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total":     len(proposals),
		"proposals": proposals,
	})
}
