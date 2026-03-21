package controllers

import (
	"logired/src/internal/cars/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetCarByIdController struct {
	getCarById *application.GetCarById
}

func NewGetCarByIdController(get *application.GetCarById) *GetCarByIdController {
	return &GetCarByIdController{
		getCarById: get,
	}
}

func (ctrl *GetCarByIdController) GetById(c *gin.Context) {
	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}
	idDriver := userIDInterface.(int32)

	idStr := c.Param("id")
	idCar, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	car, err := ctrl.getCarById.Execute(int32(idCar), idDriver)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"car": car})
}