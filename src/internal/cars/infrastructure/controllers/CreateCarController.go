package controllers

import (
	"logired/src/internal/cars/application"
	"logired/src/internal/cars/domain/entities"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	storageService "logired/src/core/services/storage"
)

type CreateCarController struct {
	createCar *application.CreateCar
	storage   storageService.StorageService 
}

func NewCreateCarController(create *application.CreateCar, storage storageService.StorageService) *CreateCarController {
	return &CreateCarController{
		createCar: create,
		storage:   storage, 
	}
}

// Create godoc
// @Summary      Crear un nuevo vehículo
// @Description  Un conductor registra su carro
// @Tags         cars
// @Accept       multipart/form-data
// @Produce      json
// @Param        car_registration  formData  string  true   "Placa"
// @Param        brand             formData  string  true   "Marca"
// @Param        model             formData  string  true   "Modelo"
// @Param        color             formData  string  true   "Color"
// @Param        max_capacity      formData  integer true   "Capacidad máxima"
// @Param        frontview_image   formData  file    false  "Imagen frontal"
// @Param        backview_image    formData  file    false  "Imagen trasera"
// @Param        leftview_image    formData  file    false  "Imagen lado izquierdo"
// @Param        rightview_image   formData  file    false  "Imagen lado derecho"
// @Param        space_image       formData  file    false  "Imagen del espacio de carga"
// @Param        plates_image      formData  file    false  "Imagen de placas"
// @Security     Bearer
// @Success      201 {object} dto.CarResponse
// @Failure      400 {object} map[string]string
// @Failure      401 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /cars [post]
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

	if err := c.Request.ParseMultipartForm(15 << 20); err != nil {
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

	// Se pasa ctrl.storage como tercer parámetro
	frontView, err := saveCarImage(c, "frontview_image", ctrl.storage)
	if err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return }

	backView, err := saveCarImage(c, "backview_image", ctrl.storage)
	if err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return }

	leftView, err := saveCarImage(c, "leftview_image", ctrl.storage)
	if err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return }

	rightView, err := saveCarImage(c, "rightview_image", ctrl.storage)
	if err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return }

	spaces, err := saveCarImage(c, "space_image", ctrl.storage)
	if err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return }

	plates, err := saveCarImage(c, "plates_image", ctrl.storage)
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
		LeftViewImage:   leftView,
		RightViewImage:  rightView,
		PlatesImage:     plates,
		SpacesImage:     spaces,
	}

	if err := ctrl.createCar.Execute(car); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Carro creado correctamente", "car": car})
}