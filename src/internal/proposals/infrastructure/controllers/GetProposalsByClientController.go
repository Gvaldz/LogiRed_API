package controllers

import (
	"logired/src/internal/proposals/application"
	"logired/src/internal/proposals/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetProposalsByClientController struct {
	uc *application.GetProposalsByClient
}

func NewGetProposalsByClientController(uc *application.GetProposalsByClient) *GetProposalsByClientController {
	return &GetProposalsByClientController{uc: uc}
}

// GetMyProposalsAsClient godoc
// @Summary      Obtener todas las propuestas recibidas por el cliente autenticado
// @Description  Devuelve todas las propuestas de todos los viajes creados por el cliente logueado.
// @Tags         proposals
// @Security     Bearer
// @Produce      json
// @Success      200 {object} map[string]interface{} "lista de propuestas"
// @Failure      401 {object} map[string]string "no autenticado"
// @Failure      500 {object} map[string]string "error interno"
// @Router       /proposals/client [get]
func (ctrl *GetProposalsByClientController) GetMyProposalsAsClient(c *gin.Context) {
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
			"message":   "No tienes propuestas en tus viajes",
			"proposals": []entities.Proposal{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total":     len(proposals),
		"proposals": proposals,
	})
}
