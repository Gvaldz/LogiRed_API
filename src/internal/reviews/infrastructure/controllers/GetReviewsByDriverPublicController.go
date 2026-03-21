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