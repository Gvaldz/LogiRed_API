package repositories

import (
	"database/sql"
	"logired/src/internal/payments/domain/entities"
)

type PaymentRepo struct {
	db *sql.DB
}

func NewPaymentRepo(db *sql.DB) *PaymentRepo {
	return &PaymentRepo{db: db}
}

func (r *PaymentRepo) CreatePayment(p entities.Payment) error {
	_, err := r.db.Exec(`
		INSERT INTO payments (id_ride, id_client, amount, currency, status, stripe_payment_intent)
		VALUES ($1, $2, $3, $4, $5, $6)`,
		p.IdRide, p.IdClient, p.Amount, p.Currency, p.Status, p.StripePaymentIntent,
	)
	return err
}

func (r *PaymentRepo) GetPaymentByRide(rideId int32) (*entities.Payment, error) {
	p := &entities.Payment{}
	err := r.db.QueryRow(`
		SELECT id_payment, id_ride, id_client, amount, currency, status, stripe_payment_intent
		FROM payments WHERE id_ride = $1`, rideId,
	).Scan(&p.IdPayment, &p.IdRide, &p.IdClient, &p.Amount, &p.Currency, &p.Status, &p.StripePaymentIntent)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (r *PaymentRepo) UpdatePaymentStatus(paymentIntentId string, status string) error {
	_, err := r.db.Exec(`
		UPDATE payments SET status = $1 WHERE stripe_payment_intent = $2`,
		status, paymentIntentId,
	)
	return err
}