package controllers

import (
	"logired/src/internal/reviews/application"
	"logired/src/internal/reviews/domain/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateReviewController struct {
	updateReview *application.UpdateReview
}

func NewUpdateReviewController(update *application.UpdateReview) *UpdateReviewController {
	return &UpdateReviewController{updateReview: update}
}

// Update godoc
// @Summary      Actualizar una reseña
// @Description  El pasajero propietario puede editar su reseña (solo texto y rating)
// @Tags         reviews
// @Accept       json
// @Produce      json
// @Param        id path int true "ID de la reseña"
// @Param        body body map[string]interface{} true "Datos a actualizar" Example({"review":"Buen servicio","rating":4})
// @Security     Bearer
// @Success      200 {object} map[string]string "mensaje de éxito"
// @Failure      400 {object} map[string]string "error en la solicitud"
// @Failure      401 {object} map[string]string "no autenticado"
// @Failure      404 {object} map[string]string "reseña no encontrada"
// @Failure      500 {object} map[string]string "error interno"
// @Router       /reviews/{id} [put]
func (ctrl *UpdateReviewController) Update(c *gin.Context) {
	idParam := c.Param("id")
	idReview, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}
	passangerID := userIDInterface.(int32)

	var req struct {
		Review string  `json:"review"`
		Rating float32 `json:"rating"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	review := entities.Review{
		IdReview:    int32(idReview),
		Review:      req.Review,
		Rating:      req.Rating,
		IdPassanger: passangerID,
	}

	if err := ctrl.updateReview.Execute(review); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reseña actualizada correctamente"})
}
