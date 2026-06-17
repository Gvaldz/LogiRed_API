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
// @Description  Cambia el estado de una propuesta (1 = aceptada, 3 = rechazada)
// @Tags         proposals
// @Accept       json
// @Produce      json
// @Param        id path int true "ID de la propuesta"
// @Param        body body map[string]interface{} true "Nuevo estado" Example({"idstatus": 1})
// @Security     ApiKeyAuth
// @Success      200 {object} map[string]string "mensaje de éxito"
// @Failure      400 {object} map[string]string "error en la solicitud"
// @Failure      401 {object} map[string]string "no autenticado"
// @Failure      500 {object} map[string]string "error interno"
// @Router       /proposals/{id}/accept [put]
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

// AcceptProposalRequest representa la solicitud para aceptar/rechazar una propuesta
type AcceptProposalRequest struct {
	IdStatus int32 `json:"idstatus" example:"1" enums:"1,3"` // 1=aceptada, 3=rechazada
}
