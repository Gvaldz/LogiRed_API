package controllers

import (
	"logired/src/internal/reviews/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetReviewsByDriverPublicController struct {
	getReviewsByDriver *application.GetReviewsByDriver
}

func NewGetReviewsByDriverPublicController(get *application.GetReviewsByDriver) *GetReviewsByDriverPublicController {
	return &GetReviewsByDriverPublicController{getReviewsByDriver: get}
}

// Handle godoc
// @Summary      Obtener reseñas de un conductor por ID (público)
// @Description  Devuelve las reseñas de un conductor específico (sin autenticación)
// @Tags         reviews
// @Produce      json
// @Param        id path int true "ID del conductor"
// @Success      200 {object} map[string]interface{} "lista de reseñas"
// @Failure      400 {object} map[string]string "ID inválido"
// @Failure      500 {object} map[string]string "error interno"
// @Router       /reviews/driver/{id} [get]

func (ctrl *GetReviewsByDriverPublicController) Handle(c *gin.Context) {
	idParam := c.Param("id")
	driverID, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de conductor inválido"})
		return
	}

	reviews, err := ctrl.getReviewsByDriver.Execute(int32(driverID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reviews": reviews})
}