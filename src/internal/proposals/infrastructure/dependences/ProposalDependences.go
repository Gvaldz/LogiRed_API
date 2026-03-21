package dependencies

import (
	"database/sql"
	"logired/src/internal/proposals/application"
	"logired/src/internal/proposals/infrastructure/controllers"
	"logired/src/internal/proposals/infrastructure/repositories"
	"logired/src/internal/proposals/infrastructure/routes"
	"github.com/gin-gonic/gin"
)

type ProposalDependencies struct {
	DB             *sql.DB
	AuthMiddleware gin.HandlerFunc
}

func NewProposalDependencies(db *sql.DB, authMiddleware gin.HandlerFunc) *ProposalDependencies {
	return &ProposalDependencies{
		DB:             db,
		AuthMiddleware: authMiddleware,
	}
}

func (d *ProposalDependencies) GetRoutes() *routes.ProposalRoutes {
	proposalRepo := repositories.NewProposalRepo(d.DB)

	createProposalUseCase := application.NewCreateProposal(proposalRepo)
	acceptProposalUseCase := application.NewAcceptProposal(proposalRepo)
	deleteProposalUseCase := application.NewDeleteProposal(proposalRepo)

	createProposalController := controllers.NewCreateProposalController(createProposalUseCase)
	acceptProposalController := controllers.NewAcceptProposalController(acceptProposalUseCase)
	deleteProposalController := controllers.NewDeleteProposalController(deleteProposalUseCase)

	return routes.NewProposalRoutes(
		createProposalController,
		acceptProposalController,
		deleteProposalController,
		d.AuthMiddleware,
	)
}