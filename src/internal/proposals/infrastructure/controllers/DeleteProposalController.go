package controllers

import (
	"logired/src/internal/proposals/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// DeleteProposalController maneja la eliminación de propuestas
type DeleteProposalController struct {
	deleteProposal *application.DeleteProposal
}

func NewDeleteProposalController(delete *application.DeleteProposal) *DeleteProposalController {
	return &DeleteProposalController{deleteProposal: delete}
}

// Delete godoc
// @Summary      Eliminar una propuesta
// @Description  Solo el conductor propietario puede eliminar su propuesta
// @Tags         proposals
// @Produce      json
// @Param        id path int true "ID de la propuesta"
// @Security     Bearer
// @Success      200 {object} map[string]string "mensaje de éxito"
// @Failure      400 {object} map[string]string "ID inválido"
// @Failure      401 {object} map[string]string "no autenticado"
// @Failure      404 {object} map[string]string "propuesta no encontrada"
// @Failure      500 {object} map[string]string "error interno"
// @Router       /proposals/{id} [delete]
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