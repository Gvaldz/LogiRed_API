package controllers

import (
	"logired/src/internal/proposals/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteProposalController struct {
	deleteProposal *application.DeleteProposal
}

func NewDeleteProposalController(delete *application.DeleteProposal) *DeleteProposalController {
	return &DeleteProposalController{deleteProposal: delete}
}

func (ctrl *DeleteProposalController) Delete(c *gin.Context) {
	idParam := c.Param("id")
	idProposal, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}
	driverID := userIDInterface.(int32)

	if err := ctrl.deleteProposal.Execute(int32(idProposal), driverID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Propuesta eliminada correctamente"})
}