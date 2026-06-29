package controllers

import (
	"logired/src/internal/proposals/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AcceptProposalController maneja la aceptación o rechazo de una propuesta
type AcceptProposalController struct {
	acceptProposal *application.AcceptProposal
}

func NewAcceptProposalController(accept *application.AcceptProposal) *AcceptProposalController {
	return &AcceptProposalController{acceptProposal: accept}
}

// Accept godoc
// @Summary      Aceptar o rechazar una propuesta
// @Tags         proposals
// @Accept       json
// @Produce      json
// @Param        id path int true "ID de la propuesta"
// @Param        body body AcceptProposalRequest true "Nuevo estado"
// @Security     Bearer
// @Success      200 {object} map[string]string
// @Failure      400 {object} map[string]string
// @Router       /proposals/{id}/accept [put]
func (ctrl *AcceptProposalController) Accept(c *gin.Context) {
	idProposal, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}
	userID := userIDInterface.(int32)

	var body struct {
		IdStatus int32 `json:"idstatus" binding:"required"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "idstatus es requerido"})
		return
	}

	if body.IdStatus != 1 && body.IdStatus != 3 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Estado inválido. Use 1 (aceptada) o 3 (rechazada)"})
		return
	}

	if err := ctrl.acceptProposal.Execute(int32(idProposal), body.IdStatus, userID); err != nil {
		if err.Error() == "solo el cliente dueño del viaje puede aceptar propuestas" {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	message := "Propuesta aceptada correctamente"
	if body.IdStatus == 3 {
		message = "Propuesta rechazada correctamente"
	}

	c.JSON(http.StatusOK, gin.H{"message": message})
}
