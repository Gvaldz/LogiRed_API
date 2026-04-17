package dependencies

import (
	"database/sql"
	"logired/src/internal/payments/application"
	"logired/src/internal/payments/infrastructure/controllers"
	"logired/src/internal/payments/infrastructure/repositories"
	"logired/src/internal/payments/infrastructure/routes"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v76"
)

type PaymentDependencies struct {
	DB             *sql.DB
	AuthMiddleware gin.HandlerFunc
	StripeKey      string
}

func NewPaymentDependencies(db *sql.DB, authMiddleware gin.HandlerFunc, stripeKey string) *PaymentDependencies {
	stripe.Key = stripeKey
	return &PaymentDependencies{
		DB:             db,
		AuthMiddleware: authMiddleware,
		StripeKey:      stripeKey,
	}
}

func (d *PaymentDependencies) GetRoutes() *routes.PaymentRoutes {
	repo := repositories.NewPaymentRepo(d.DB)
	useCase := application.NewCreatePaymentIntent(repo)
	controller := controllers.NewPaymentController(useCase)
	return routes.NewPaymentRoutes(controller, d.AuthMiddleware)
}
