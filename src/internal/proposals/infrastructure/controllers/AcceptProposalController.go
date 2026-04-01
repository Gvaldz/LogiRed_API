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
	idProposal, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

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

	if err := ctrl.acceptProposal.Execute(int32(idProposal), body.IdStatus); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	message := "Propuesta aceptada correctamente"
	if body.IdStatus == 3 {
		message = "Propuesta rechazada correctamente"
	}

	c.JSON(http.StatusOK, gin.H{"message": message})
}