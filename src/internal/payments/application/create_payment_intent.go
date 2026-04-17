package application

import (
	"logired/src/internal/payments/domain"
	"logired/src/internal/payments/domain/entities"

	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/paymentintent"
)

type CreatePaymentIntent struct {
	repo domain.IPayment
}

func NewCreatePaymentIntent(repo domain.IPayment) *CreatePaymentIntent {
	return &CreatePaymentIntent{repo: repo}
}

type PaymentIntentResult struct {
	ClientSecret    string `json:"client_secret"`
	PaymentIntentId string `json:"payment_intent_id"`
}

func (uc *CreatePaymentIntent) Execute(rideId int32, clientId int32, amount float64) (*PaymentIntentResult, error) {
	amountCents := int64(amount * 100)

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(amountCents),
		Currency: stripe.String("mxn"),
		Metadata: map[string]string{
			"ride_id":   string(rune(rideId)),
			"client_id": string(rune(clientId)),
		},
	}

	pi, err := paymentintent.New(params)
	if err != nil {
		return nil, err
	}

	payment := entities.Payment{
		IdRide:              rideId,
		IdClient:            clientId,
		Amount:              amount,
		Currency:            "mxn",
		Status:              "pending",
		StripePaymentIntent: pi.ID,
	}

	if err := uc.repo.CreatePayment(payment); err != nil {
		return nil, err
	}

	return &PaymentIntentResult{
		ClientSecret:    pi.ClientSecret,
		PaymentIntentId: pi.ID,
	}, nil
}
