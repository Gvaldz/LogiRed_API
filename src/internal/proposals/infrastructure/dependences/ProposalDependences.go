package dependencies

import (
	"database/sql"
	notifications "logired/src/core/services/notifications"
	"logired/src/internal/proposals/application"
	"logired/src/internal/proposals/infrastructure/controllers"
	"logired/src/internal/proposals/infrastructure/repositories"
	"logired/src/internal/proposals/infrastructure/routes"
	rideRepositories "logired/src/internal/rides/infrastructure/repositories"

	"github.com/gin-gonic/gin"
)

type ProposalDependencies struct {
	DB                           *sql.DB
	AuthMiddleware               gin.HandlerFunc
	DriverOnlyApprovedMiddleware gin.HandlerFunc
	Notifier                     *notifications.NotificationService
}

func NewProposalDependencies(
	db *sql.DB,
	authMiddleware gin.HandlerFunc,
	driverOnlyApprovedMiddleware gin.HandlerFunc,
	notifier *notifications.NotificationService,
) *ProposalDependencies {
	return &ProposalDependencies{
		DB:                           db,
		AuthMiddleware:               authMiddleware,
		DriverOnlyApprovedMiddleware: driverOnlyApprovedMiddleware,
		Notifier:                     notifier,
	}
}

func (d *ProposalDependencies) GetRoutes() *routes.ProposalRoutes {
	proposalRepo := repositories.NewProposalRepo(d.DB)
	rideRepo := rideRepositories.NewRideRepo(d.DB)

	createProposalUseCase := application.NewCreateProposal(proposalRepo, rideRepo, d.Notifier)
	acceptProposalUseCase := application.NewAcceptProposal(proposalRepo, rideRepo, d.Notifier)
	getProposalByIdUseCase := application.NewGetProposalById(proposalRepo, rideRepo)
	getProposalsByRideUseCase := application.NewGetProposalsByRide(proposalRepo, rideRepo)
	deleteProposalUseCase := application.NewDeleteProposal(proposalRepo)

	createProposalController := controllers.NewCreateProposalController(createProposalUseCase)
	acceptProposalController := controllers.NewAcceptProposalController(acceptProposalUseCase)
	deleteProposalController := controllers.NewDeleteProposalController(deleteProposalUseCase)
	getProposalByIdController := controllers.NewGetProposalyIdController(getProposalByIdUseCase)
	getProposalsByRideController := controllers.NewGetProposalsByRideController(getProposalsByRideUseCase)

	return routes.NewProposalRoutes(
		createProposalController,
		acceptProposalController,
		getProposalByIdController,
		getProposalsByRideController,
		deleteProposalController,
		d.AuthMiddleware,
		d.DriverOnlyApprovedMiddleware,
	)
}
