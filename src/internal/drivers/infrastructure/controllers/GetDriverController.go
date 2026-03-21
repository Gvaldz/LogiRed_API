package controllers

import (
	"logired/src/internal/drivers/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetDriverController struct {
	getByUser *application.GetDriverByUser
	getByID   *application.GetDriverByID
}

func NewGetDriverController(getByUser *application.GetDriverByUser, getByID *application.GetDriverByID) *GetDriverController {
	return &GetDriverController{
		getByUser: getByUser,
		getByID:   getByID,
	}
}

func (ctrl *GetDriverController) GetMe(c *gin.Context) {
	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}
	userID := userIDInterface.(int32)

	driver, err := ctrl.getByUser.Execute(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"driver": driver})
}

func (ctrl *GetDriverController) GetByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	driver, err := ctrl.getByID.Execute(int32(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"driver": driver})
}