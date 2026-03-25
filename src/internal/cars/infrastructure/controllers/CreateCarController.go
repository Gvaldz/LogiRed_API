package controllers

import (
	"logired/src/internal/cars/application"
	"logired/src/internal/cars/domain/entities"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type CreateCarController struct {
	createCar *application.CreateCar
}

func NewCreateCarController(create *application.CreateCar) *CreateCarController {
	return &CreateCarController{createCar: create}
}

func (ctrl *CreateCarController) Create(c *gin.Context) {
	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}
	idDriver := userIDInterface.(int32)

	if !strings.HasPrefix(c.GetHeader("Content-Type"), "multipart/form-data") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Content-Type debe ser multipart/form-data"})
		return
	}

	if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al parsear formulario: " + err.Error()})
		return
	}

	carRegistration := c.Request.FormValue("car_registration")
	brand           := c.Request.FormValue("brand")
	model           := c.Request.FormValue("model")
	color           := c.Request.FormValue("color")
	maxCapacityStr  := c.Request.FormValue("max_capacity")

	if carRegistration == "" || brand == "" || model == "" || color == "" || maxCapacityStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Faltan campos obligatorios"})
		return
	}

	maxCapacity, err := strconv.ParseInt(maxCapacityStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "max_capacity inválido"})
		return
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
		IdDriver:        idDriver,
		CarRegistration: carRegistration,
		Brand:           brand,
		Model:           model,
		Color:           color,
		MaxCapacity:     int32(maxCapacity),
		FrontViewImage:  frontView,
		BackViewImage:   backView,
		PlatesImage:     plates,
		SpacesImage:     spaces,
	}

	if err := ctrl.createCar.Execute(car); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Carro creado correctamente", "car": car})
}