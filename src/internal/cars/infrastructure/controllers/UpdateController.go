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

type UpdateCarController struct {
	updateCar *application.UpdateCar
	storage   storageService.StorageService 
}

func NewUpdateCarController(update *application.UpdateCar, storage storageService.StorageService) *UpdateCarController {
	return &UpdateCarController{
		updateCar: update,
		storage:   storage,
	}
}
// Update godoc
// @Summary      Actualizar un vehículo
// @Description  El conductor propietario puede actualizar los datos de su carro (incluye imágenes)
// @Tags         cars
// @Accept       multipart/form-data
// @Produce      json
// @Param        id path int true "ID del carro"
// @Param        car_registration formData string false "Placa"
// @Param        brand formData string false "Marca"
// @Param        model formData string false "Modelo"
// @Param        color formData string false "Color"
// @Param        max_capacity formData integer false "Capacidad máxima"
// @Param        frontview_image formData file false "Imagen frontal"
// @Param        backview_image formData file false "Imagen trasera"
// @Param        leftview_image formData file false "Imagen lado izquierdo"
// @Param        rightview_image formData file false "Imagen lado derecho"
// @Param        space_image formData file false "Imagen del espacio de carga"
// @Param        plates_image formData file false "Imagen de placas"
// @Security     Bearer
// @Success      200 {object} map[string]interface{} "mensaje y carro actualizado"
// @Failure      400 {object} map[string]string "error en los datos"
// @Failure      401 {object} map[string]string "no autenticado"
// @Failure      404 {object} map[string]string "carro no encontrado"
// @Failure      500 {object} map[string]string "error interno"
// @Router       /cars/{id} [put]
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

	if err := c.Request.ParseMultipartForm(15 << 20); err != nil {
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
		IdCar:           int32(idCar),
		IdDriver:        idDriver,
		CarRegistration: c.Request.FormValue("car_registration"),
		Brand:           c.Request.FormValue("brand"),
		Model:           c.Request.FormValue("model"),
		Color:           c.Request.FormValue("color"),
		MaxCapacity:     maxCapacity,
		FrontViewImage:  frontView,
		BackViewImage:   backView,
		LeftViewImage:   leftView,
		RightViewImage:  rightView,
		PlatesImage:     plates,
		SpacesImage:     spaces,
	}

	if err := ctrl.updateCar.Execute(car); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Carro actualizado correctamente", "car": car})
}