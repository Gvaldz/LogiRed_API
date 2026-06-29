package controllers

import (
	"logired/src/internal/reviews/application"
	"logired/src/internal/reviews/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateReviewController struct {
	createReview *application.CreateReview
}

func NewCreateReviewController(create *application.CreateReview) *CreateReviewController {
	return &CreateReviewController{createReview: create}
}

// Create godoc
// @Summary      Crear una reseña
// @Description  Un pasajero crea una reseña para un conductor
// @Tags         reviews
// @Accept       json
// @Produce      json
// @Param        body body map[string]interface{} true "Datos de la reseña" Example({"review":"Excelente","rating":5,"iduser":14})
// @Security     Bearer
// @Success      201 {object} map[string]interface{} "mensaje y reseña creada"
// @Failure      400 {object} map[string]string "error en los datos"
// @Failure      401 {object} map[string]string "no autenticado"
// @Failure      500 {object} map[string]string "error interno"
// @Router       /reviews [post]
func (ctrl *CreateReviewController) Create(c *gin.Context) {
	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}
	passangerID := userIDInterface.(int32)

	var req struct {
		Review   string  `json:"review"`
		Rating   float32 `json:"rating"`
		IdDriver int32   `json:"iduser"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	review := entities.Review{
		Review:      req.Review,
		Rating:      req.Rating,
		IdDriver:    req.IdDriver,
		IdPassanger: passangerID,
	}

	if err := ctrl.createReview.Execute(review); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Reseña creada correctamente", "review": review})
}
