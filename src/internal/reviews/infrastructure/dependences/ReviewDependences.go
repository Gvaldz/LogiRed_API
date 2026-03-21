package dependencies

import (
	"database/sql"
	"logired/src/internal/reviews/application"
	"logired/src/internal/reviews/infrastructure/controllers"
	"logired/src/internal/reviews/infrastructure/repositories"
	"logired/src/internal/reviews/infrastructure/routes"

	"github.com/gin-gonic/gin"
)

type ReviewDependencies struct {
	DB             *sql.DB
	AuthMiddleware gin.HandlerFunc
}

func NewReviewDependencies(db *sql.DB, authMiddleware gin.HandlerFunc) *ReviewDependencies {
	return &ReviewDependencies{
		DB:             db,
		AuthMiddleware: authMiddleware,
	}
}

func (d *ReviewDependencies) GetRoutes() *routes.ReviewRoutes {
	reviewRepo := repositories.NewReviewRepo(d.DB)

	createReviewUseCase := application.NewCreateReview(reviewRepo)
	getReviewsByDriverUseCase := application.NewGetReviewsByDriver(reviewRepo)
	getReviewsByPassangerUseCase := application.NewGetReviewsByPassanger(reviewRepo)
	updateReviewUseCase := application.NewUpdateReview(reviewRepo)

	createReviewController := controllers.NewCreateReviewController(createReviewUseCase)
	getReviewsByDriverController := controllers.NewGetReviewsByDriverController(getReviewsByDriverUseCase)
	getReviewsByPassangerController := controllers.NewGetReviewsByPassangerController(getReviewsByPassangerUseCase)
	getReviewsByDriverPublicController := controllers.NewGetReviewsByDriverPublicController(getReviewsByDriverUseCase)
	updateReviewController := controllers.NewUpdateReviewController(updateReviewUseCase)

	return routes.NewReviewRoutes(
		createReviewController,
		getReviewsByDriverController,
		getReviewsByPassangerController,
		getReviewsByDriverPublicController,
		updateReviewController,
		d.AuthMiddleware,
	)
}