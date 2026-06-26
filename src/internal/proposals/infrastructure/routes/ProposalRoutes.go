package routes

import (
	"logired/src/internal/proposals/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

type ProposalRoutes struct {
	createProposalController      *controllers.CreateProposalController
	acceptProposalController      *controllers.AcceptProposalController
	getProposalByIdController     *controllers.GetProposalByIdController
	getProposalsByRideController  *controllers.GetProposalsByRideController
	deleteProposalController      *controllers.DeleteProposalController
	authMiddleware                gin.HandlerFunc
	driverOnlyApprovedMiddleware  gin.HandlerFunc
}

func NewProposalRoutes(
	create *controllers.CreateProposalController,
	accept *controllers.AcceptProposalController,
	getById *controllers.GetProposalByIdController,
	getProposalsByRide *controllers.GetProposalsByRideController,
	delete *controllers.DeleteProposalController,
	authMiddleware gin.HandlerFunc,
	driverOnlyApprovedMiddleware gin.HandlerFunc,
) *ProposalRoutes {
	return &ProposalRoutes{
		createProposalController:     create,
		acceptProposalController:     accept,
		getProposalByIdController:    getById,
		getProposalsByRideController: getProposalsByRide,
		deleteProposalController:     delete,
		authMiddleware:               authMiddleware,
		driverOnlyApprovedMiddleware: driverOnlyApprovedMiddleware,
	}
}

func (r *ProposalRoutes) AttachRoutes(router *gin.Engine) {
	// TODAS las rutas de propuestas requieren que sea conductor aprobado
	proposalsGroup := router.Group("/proposals")
	proposalsGroup.Use(r.authMiddleware)
	proposalsGroup.Use(r.driverOnlyApprovedMiddleware)

	proposalsGroup.POST("", r.createProposalController.Create)
	proposalsGroup.GET("/:id", r.getProposalByIdController.GetById)
	proposalsGroup.GET("/ride/:idride", r.getProposalsByRideController.GetByRide)
	proposalsGroup.PUT("/:id/accept", r.acceptProposalController.Accept)
	proposalsGroup.DELETE("/:id", r.deleteProposalController.Delete)
}
