package routes

import (
	"logired/src/internal/drivers/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

type DriverRoutes struct {
	createController *controllers.CreateDriverController
	getController    *controllers.GetDriverController
	getAllController *controllers.GetAllDriversController
	deleteController *controllers.DeleteDriverController
	authMiddleware   gin.HandlerFunc
}

func NewDriverRoutes(
	create *controllers.CreateDriverController,
	get *controllers.GetDriverController,
	getAll *controllers.GetAllDriversController,
	delete *controllers.DeleteDriverController,
	authMiddleware gin.HandlerFunc,
) *DriverRoutes {
	return &DriverRoutes{
		createController: create,
		getController:    get,
		getAllController: getAll,
		deleteController: delete,
		authMiddleware:   authMiddleware,
	}
}

func (r *DriverRoutes) AttachRoutes(router *gin.Engine) {
	driversGroup := router.Group("/drivers")
	driversGroup.Use(r.authMiddleware)

	driversGroup.POST("", r.createController.Create)
	driversGroup.GET("/me", r.getController.GetMe)
	driversGroup.GET("/:id", r.getController.GetByID)
	driversGroup.GET("", r.getAllController.GetAll)
	driversGroup.DELETE("/me", r.deleteController.Delete)
}