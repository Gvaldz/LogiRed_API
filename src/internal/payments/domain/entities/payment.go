package entities

type Payment struct {
	IdPayment           int32   `json:"id_payment"`
	IdRide              int32   `json:"id_ride"`
	IdClient            int32   `json:"id_client"`
	Amount              float64 `json:"amount"`
	Currency            string  `json:"currency"`
	Status              string  `json:"status"`
	StripePaymentIntent string  `json:"stripe_payment_intent"`
}
