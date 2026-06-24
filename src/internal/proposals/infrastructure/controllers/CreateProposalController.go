package controllers

import (
	"logired/src/internal/proposals/application"
	"logired/src/internal/proposals/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateProposalController maneja la creación de propuestas
type CreateProposalController struct {
	createProposal *application.CreateProposal
}

func NewCreateProposalController(create *application.CreateProposal) *CreateProposalController {
	return &CreateProposalController{createProposal: create}
}

// Create godoc
// @Summary      Crear una propuesta
// @Tags         proposals
// @Accept       json
// @Produce      json
// @Param        body body CreateProposalRequest true "Datos de la propuesta"
// @Security     ApiKeyAuth
// @Success      201 {object} map[string]interface{}
// @Failure      400 {object} map[string]string
// @Router       /proposals [post]
func (ctrl *CreateProposalController) Create(c *gin.Context) {
	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}
	driverID := userIDInterface.(int32)

	var req CreateProposalRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	proposal := entities.Proposal{
		Price:    req.Price,
		IdDriver: driverID,
		IdRide:   req.IdRide,
		Comment:  req.Comment,
		IdCar:    req.IdCar,
		IdStatus: 2, 
	}

	if err := ctrl.createProposal.Execute(proposal); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":  "Propuesta creada correctamente",
		"proposal": proposal,
	})
}

