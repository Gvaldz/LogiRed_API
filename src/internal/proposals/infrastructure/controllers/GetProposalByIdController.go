package controllers

import (
	"logired/src/internal/proposals/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetProposalByIdController maneja la obtención de una propuesta por su ID
type GetProposalByIdController struct {
	getProposalyId *application.GetProposalById
}

func NewGetProposalyIdController(get *application.GetProposalById) *GetProposalByIdController {
	return &GetProposalByIdController{getProposalyId: get}
}

// GetById godoc
// @Summary      Obtener una propuesta por ID
// @Tags         proposals
// @Produce      json
// @Param        id path int true "ID de la propuesta"
// @Success      200 {object} map[string]interface{} "datos de la propuesta"
// @Failure      400 {object} map[string]string "ID inválido"
// @Failure      404 {object} map[string]string "propuesta no encontrada"
// @Router       /proposals/{id} [get]
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

// ProposalResponse representa la estructura de una propuesta en la respuesta
type ProposalResponse struct {
	IdProposal int32   `json:"idproposal" example:"1"`
	Price      float32 `json:"price" example:"150.50"`
	Comment    string  `json:"comment" example:"Puedo llevar tu paquete"`
	IdDriver   int32   `json:"iddriver" example:"14"`
	IdRide     int32   `json:"idride" example:"1"`
	IdStatus   int32   `json:"idproposalstatus" example:"2"`
	IdCar      int32   `json:"idcar" example:"7"`
}