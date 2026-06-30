package routes

import (
	"logired/src/internal/proposals/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

type ProposalRoutes struct {
	createProposalController        *controllers.CreateProposalController
	acceptProposalController        *controllers.AcceptProposalController
	getProposalByIdController       *controllers.GetProposalByIdController
	getProposalsByRideController    *controllers.GetProposalsByRideController
	getProposalsByDriverController  *controllers.GetProposalsByDriverController
	getProposalsByClientController  *controllers.GetProposalsByClientController
	deleteProposalController        *controllers.DeleteProposalController
	authMiddleware                  gin.HandlerFunc
	driverOnlyApprovedMiddleware    gin.HandlerFunc
}

func NewProposalRoutes(
	create *controllers.CreateProposalController,
	accept *controllers.AcceptProposalController,
	getById *controllers.GetProposalByIdController,
	getProposalsByRide *controllers.GetProposalsByRideController,
	getProposalsByDriver *controllers.GetProposalsByDriverController,
	getProposalsByClient *controllers.GetProposalsByClientController,
	delete *controllers.DeleteProposalController,
	authMiddleware gin.HandlerFunc,
	driverOnlyApprovedMiddleware gin.HandlerFunc,
) *ProposalRoutes {
	return &ProposalRoutes{
		createProposalController:       create,
		acceptProposalController:       accept,
		getProposalByIdController:      getById,
		getProposalsByRideController:   getProposalsByRide,
		getProposalsByDriverController: getProposalsByDriver,
		getProposalsByClientController: getProposalsByClient,
		deleteProposalController:       delete,
		authMiddleware:                 authMiddleware,
		driverOnlyApprovedMiddleware:   driverOnlyApprovedMiddleware,
	}
}

func (r *ProposalRoutes) AttachRoutes(router *gin.Engine) {
	// Solo conductores aprobados pueden crear, eliminar y ver sus propias propuestas
	driverGroup := router.Group("/proposals")
	driverGroup.Use(r.authMiddleware)
	driverGroup.Use(r.driverOnlyApprovedMiddleware)
	driverGroup.POST("", r.createProposalController.Create)
	driverGroup.DELETE("/:id", r.deleteProposalController.Delete)
	driverGroup.GET("/driver", r.getProposalsByDriverController.GetMyProposals)

	// Cualquier usuario autenticado puede ver y aceptar propuestas
	authGroup := router.Group("/proposals")
	authGroup.Use(r.authMiddleware)
	authGroup.GET("/:id", r.getProposalByIdController.GetById)
	authGroup.GET("/ride/:idride", r.getProposalsByRideController.GetByRide)
	authGroup.GET("/client", r.getProposalsByClientController.GetMyProposalsAsClient)
	authGroup.PUT("/:id/accept", r.acceptProposalController.Accept)
}
