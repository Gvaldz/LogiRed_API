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

// GetById godoc
// @Summary      Obtener una propuesta por ID con datos del conductor y su vehículo
// @Tags         proposals
// @Security     Bearer
// @Produce      json
// @Param        id path int true "ID de la propuesta"
// @Success      200 {object} map[string]interface{} "datos de la propuesta"
// @Failure      400 {object} map[string]string "ID inválido"
// @Failure      401 {object} map[string]string "no autenticado"
// @Failure      404 {object} map[string]string "propuesta no encontrada"
// @Router       /proposals/{id} [get]
func (ctrl *GetProposalByIdController) GetById(c *gin.Context) {
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
	userID := userIDInterface.(int32)

	userTypeInterface, _ := c.Get("userType")
	userType, _ := userTypeInterface.(int)

	detail, err := ctrl.getProposalyId.Execute(int32(idProposal), userID, userType)
	if err != nil {
		if err.Error() == "no tienes permiso para ver esta propuesta" {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"proposal": detail})
}
