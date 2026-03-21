package routes

import (
	"logired/src/internal/drivers/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

type DriverRoutes struct {
	createController *controllers.CreateDriverController
	getController    *controllers.GetDriverController
	getAllController *controllers.GetAllDriversController
	updateController *controllers.UpdateDriverController
	deleteController *controllers.DeleteDriverController
	authMiddleware   gin.HandlerFunc
}

func NewDriverRoutes(
	create *controllers.CreateDriverController,
	get *controllers.GetDriverController,
	getAll *controllers.GetAllDriversController,
	update *controllers.UpdateDriverController,
	delete *controllers.DeleteDriverController,
	authMiddleware gin.HandlerFunc,
) *DriverRoutes {
	return &DriverRoutes{
		createController: create,
		getController:    get,
		getAllController: getAll,
		updateController: update,
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
	driversGroup.PUT("/me", r.updateController.Update)
	driversGroup.DELETE("/me", r.deleteController.Delete)
}