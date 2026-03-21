package routes

import (
	"logired/src/internal/cars/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

type CarRoutes struct {
	createCarController  *controllers.CreateCarController
	updateCarController  *controllers.UpdateCarController
	getCarsByDriverController *controllers.GetCarsByDriverController
	getCarByIdController *controllers.GetCarByIdController
	deleteCarController  *controllers.DeleteCarController
	authMiddleware       gin.HandlerFunc
}

func NewCarRoutes(
	createCarController *controllers.CreateCarController,
	updateCarController *controllers.UpdateCarController,
	getCarsByDriverController *controllers.GetCarsByDriverController,
	getCarByIdController *controllers.GetCarByIdController,
	deleteCarController *controllers.DeleteCarController,
	authMiddleware gin.HandlerFunc,
) *CarRoutes {
	return &CarRoutes{
		createCarController:  createCarController,
		updateCarController:  updateCarController,
		getCarsByDriverController: getCarsByDriverController,
		getCarByIdController: getCarByIdController,
		deleteCarController:  deleteCarController,
		authMiddleware:       authMiddleware,
	}
}

func (r *CarRoutes) AttachRoutes(router *gin.Engine) {
	carsGroup := router.Group("/cars")
	carsGroup.Use(r.authMiddleware)

	carsGroup.POST("", r.createCarController.Create)
	carsGroup.GET("", r.getCarsByDriverController.GetByDriver)
	carsGroup.GET("/:id", r.getCarByIdController.GetById)
	carsGroup.PUT("/:id", r.updateCarController.Update)
	carsGroup.DELETE("/:id", r.deleteCarController.Delete)
}