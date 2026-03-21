package routes

import (
	"logired/src/internal/proposals/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

type ProposalRoutes struct {
	createProposalController *controllers.CreateProposalController
	acceptProposalController *controllers.AcceptProposalController
	deleteProposalController *controllers.DeleteProposalController
	authMiddleware           gin.HandlerFunc
}

func NewProposalRoutes(
	create *controllers.CreateProposalController,
	accept *controllers.AcceptProposalController,
	delete *controllers.DeleteProposalController,
	authMiddleware gin.HandlerFunc,
) *ProposalRoutes {
	return &ProposalRoutes{
		createProposalController: create,
		acceptProposalController: accept,
		deleteProposalController: delete,
		authMiddleware:           authMiddleware,
	}
}

func (r *ProposalRoutes) AttachRoutes(router *gin.Engine) {
	proposalsGroup := router.Group("/proposals")
	proposalsGroup.Use(r.authMiddleware)

	proposalsGroup.POST("", r.createProposalController.Create)
	proposalsGroup.PUT("/:id/accept", r.acceptProposalController.Accept)
	proposalsGroup.DELETE("/:id", r.deleteProposalController.Delete)
}