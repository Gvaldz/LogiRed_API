package controllers

import (
	"logired/src/internal/proposals/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AcceptProposalController struct {
	acceptProposal *application.AcceptProposal
}

func NewAcceptProposalController(accept *application.AcceptProposal) *AcceptProposalController {
	return &AcceptProposalController{acceptProposal: accept}
}

func (ctrl *AcceptProposalController) Accept(c *gin.Context) {
	idParam := c.Param("id")
	idProposal, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := ctrl.acceptProposal.Execute(int32(idProposal)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Propuesta aceptada correctamente"})
}