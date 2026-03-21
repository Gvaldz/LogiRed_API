package controllers

import (
	"logired/src/internal/drivers/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllDriversController struct {
	getAll *application.GetAllDrivers
}

func NewGetAllDriversController(getAll *application.GetAllDrivers) *GetAllDriversController {
	return &GetAllDriversController{getAll: getAll}
}

func (ctrl *GetAllDriversController) GetAll(c *gin.Context) {
	drivers, err := ctrl.getAll.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"drivers": drivers})
}