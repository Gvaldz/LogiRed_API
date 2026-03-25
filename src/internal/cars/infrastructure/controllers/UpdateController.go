package controllers

import (
	"logired/src/internal/cars/application"
	"logired/src/internal/cars/domain/entities"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type UpdateCarController struct {
	updateCar *application.UpdateCar
}

func NewUpdateCarController(update *application.UpdateCar) *UpdateCarController {
	return &UpdateCarController{updateCar: update}
}

func (ctrl *UpdateCarController) Update(c *gin.Context) {
	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}
	idDriver := userIDInterface.(int32)

	idCar, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if !strings.HasPrefix(c.GetHeader("Content-Type"), "multipart/form-data") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Content-Type debe ser multipart/form-data"})
		return
	}

	if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al parsear formulario: " + err.Error()})
		return
	}

	var maxCapacity int32
	if v := c.Request.FormValue("max_capacity"); v != "" {
		parsed, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "max_capacity inválido"})
			return
		}
		maxCapacity = int32(parsed)
	}

	frontView, err := saveCarImage(c, "frontview_image")
	if err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return }

	backView, err := saveCarImage(c, "backview_image")
	if err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return }

	plates, err := saveCarImage(c, "plates_image")
	if err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return }

	spaces, err := saveCarImage(c, "space_image")
	if err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return }

	car := entities.Car{
		IdCar:           int32(idCar),
		IdDriver:        idDriver,
		CarRegistration: c.Request.FormValue("car_registration"),
		Brand:           c.Request.FormValue("brand"),
		Model:           c.Request.FormValue("model"),
		Color:           c.Request.FormValue("color"),
		MaxCapacity:     maxCapacity,
		FrontViewImage:  frontView,
		BackViewImage:   backView,
		PlatesImage:     plates,
		SpacesImage:     spaces,
	}

	if err := ctrl.updateCar.Execute(car); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Carro actualizado correctamente", "car": car})
}