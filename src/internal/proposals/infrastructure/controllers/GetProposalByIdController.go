package controllers

import (
	"logired/src/internal/proposals/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetProposalByIdController struct {
	getProposalyId *application.GetProposalById
}

func NewGetProposalyIdController(get *application.GetProposalById) *GetProposalByIdController {
	return &GetProposalByIdController{getProposalyId: get}
}

func (ctrl *GetProposalByIdController) GetById(c *gin.Context) {
	idParam := c.Param("id")
	idProposal, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	proposal, err := ctrl.getProposalyId.Execute(int32(idProposal))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"proposal": proposal})
}