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
// @Summary      Crear una nueva propuesta
// @Description  Un conductor crea una propuesta para un viaje
// @Tags         proposals
// @Accept       json
// @Produce      json
// @Param        body body map[string]interface{} true "Datos de la propuesta" Example({"price": 150.5, "id_ride": 1, "comment": "Puedo llevar", "idcar": 3})
// @Security     ApiKeyAuth
// @Success      201 {object} map[string]interface{} "mensaje y propuesta creada"
// @Failure      400 {object} map[string]string "error en los datos"
// @Failure      401 {object} map[string]string "no autenticado"
// @Failure      500 {object} map[string]string "error interno"
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
		IdStatus: 2, // Estado inicial: pendiente
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

// CreateProposalRequest define los datos necesarios para crear una propuesta
type CreateProposalRequest struct {
	Price   float32 `json:"price" example:"150.50" binding:"required"`
	IdRide  int32   `json:"id_ride" example:"1" binding:"required"`
	Comment string  `json:"comment" example:"Puedo llevar tu paquete"`
	IdCar   int32   `json:"idcar" example:"7" binding:"required"`
}
