// src/internal/tracking/infrastructure/routes/tracking_routes.go
package routes

import (
    "logired/src/internal/tracking/infrastructure/controllers"

    "github.com/gin-gonic/gin"
)

type TrackingRoutes struct {
    controller *controllers.TrackingController
}

func NewTrackingRoutes(controller *controllers.TrackingController) *TrackingRoutes {
    return &TrackingRoutes{controller: controller}
}

func (r *TrackingRoutes) AttachRoutes(router *gin.Engine) {
    ws := router.Group("/ws/rides")
    ws.GET("/:id/publish", r.controller.PublishLocation)
    ws.GET("/:id/subscribe", r.controller.SubscribeLocation)
}