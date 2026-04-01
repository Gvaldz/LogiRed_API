package routes

import (
	"logired/src/internal/proposals/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

type ProposalRoutes struct {
	createProposalController *controllers.CreateProposalController
	acceptProposalController *controllers.AcceptProposalController
	getProposalByIdController *controllers.GetProposalByIdController
	deleteProposalController *controllers.DeleteProposalController
	authMiddleware           gin.HandlerFunc
}

func NewProposalRoutes(
	create *controllers.CreateProposalController,
	accept *controllers.AcceptProposalController,
	gerById *controllers.GetProposalByIdController,
	delete *controllers.DeleteProposalController,
	authMiddleware gin.HandlerFunc,
) *ProposalRoutes {
	return &ProposalRoutes{
		createProposalController: create,
		acceptProposalController: accept,
		getProposalByIdController: gerById,
		deleteProposalController: delete,
		authMiddleware:           authMiddleware,
	}
}

func (r *ProposalRoutes) AttachRoutes(router *gin.Engine) {
	proposalsGroup := router.Group("/proposals")
	proposalsGroup.Use(r.authMiddleware)
	proposalsGroup.POST("", r.createProposalController.Create)
	proposalsGroup.GET("/:id", r.getProposalByIdController.GetById)
	proposalsGroup.PUT("/:id/accept", r.acceptProposalController.Accept)
	proposalsGroup.DELETE("/:id", r.deleteProposalController.Delete)
}