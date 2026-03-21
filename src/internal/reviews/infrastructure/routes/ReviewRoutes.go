package routes

import (
	"logired/src/internal/reviews/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

type ReviewRoutes struct {
	createReviewController          *controllers.CreateReviewController
	getReviewsByDriverController     *controllers.GetReviewsByDriverController
	getReviewsByPassangerController  *controllers.GetReviewsByPassangerController
	updateReviewController           *controllers.UpdateReviewController
	goetReviewsByDriverPublicController *controllers.GetReviewsByDriverPublicController
	authMiddleware                   gin.HandlerFunc
}

func NewReviewRoutes(
	create *controllers.CreateReviewController,
	getByDriver *controllers.GetReviewsByDriverController,
	getByPassanger *controllers.GetReviewsByPassangerController,
	getByDriverPublic *controllers.GetReviewsByDriverPublicController,
	update *controllers.UpdateReviewController,
	authMiddleware gin.HandlerFunc,
) *ReviewRoutes {
	return &ReviewRoutes{
		createReviewController:          create,
		getReviewsByDriverController:     getByDriver,
		getReviewsByPassangerController:  getByPassanger,
		updateReviewController:           update,
		authMiddleware:                   authMiddleware,
	}
}

func (r *ReviewRoutes) AttachRoutes(router *gin.Engine) {
	reviewsGroup := router.Group("/reviews")
	reviewsGroup.Use(r.authMiddleware)

	reviewsGroup.POST("", r.createReviewController.Create)
	reviewsGroup.GET("/driver/me", r.getReviewsByDriverController.GetByDriver)       
	reviewsGroup.GET("/passenger/me", r.getReviewsByPassangerController.GetByPassanger) 
	reviewsGroup.GET("/driver/:id", r.getReviewsByDriverController.GetByDriver)
	reviewsGroup.PUT("/:id", r.updateReviewController.Update)
}