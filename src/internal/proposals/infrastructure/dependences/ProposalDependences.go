package dependencies

import (
	"database/sql"
	"logired/src/internal/proposals/application"
	"logired/src/internal/proposals/infrastructure/controllers"
	"logired/src/internal/proposals/infrastructure/repositories"
	"logired/src/internal/proposals/infrastructure/routes"
	rideRepositories "logired/src/internal/rides/infrastructure/repositories"
	notifications "logired/src/internal/services/notifications"

	"github.com/gin-gonic/gin"
)

type ProposalDependencies struct {
	DB             *sql.DB
	AuthMiddleware gin.HandlerFunc
	Notifier       *notifications.NotificationService
}

func NewProposalDependencies(
	db *sql.DB,
	authMiddleware gin.HandlerFunc,
	notifier *notifications.NotificationService,
) *ProposalDependencies {
	return &ProposalDependencies{
		DB:             db,
		AuthMiddleware: authMiddleware,
		Notifier:       notifier,
	}
}

func (d *ProposalDependencies) GetRoutes() *routes.ProposalRoutes {
	proposalRepo := repositories.NewProposalRepo(d.DB)
	rideRepo     := rideRepositories.NewRideRepo(d.DB)

	createProposalUseCase := application.NewCreateProposal(proposalRepo, rideRepo, d.Notifier)
	acceptProposalUseCase := application.NewAcceptProposal(proposalRepo, rideRepo, d.Notifier)
	getProposalByIdUseCase := application.NewGetProposalById(proposalRepo)
	deleteProposalUseCase := application.NewDeleteProposal(proposalRepo)

	createProposalController := controllers.NewCreateProposalController(createProposalUseCase)
	acceptProposalController := controllers.NewAcceptProposalController(acceptProposalUseCase)
	deleteProposalController := controllers.NewDeleteProposalController(deleteProposalUseCase)
	getProposalByIdController := controllers.NewGetProposalyIdController(getProposalByIdUseCase)

	return routes.NewProposalRoutes(
		createProposalController,
		acceptProposalController,
		getProposalByIdController,
		deleteProposalController,
		d.AuthMiddleware,
	)
}