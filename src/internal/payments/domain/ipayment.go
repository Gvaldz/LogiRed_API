package domain

import "logired/src/internal/payments/domain/entities"

type IPayment interface {
	CreatePayment(payment entities.Payment) error
	GetPaymentByRide(rideId int32) (*entities.Payment, error)
	UpdatePaymentStatus(paymentIntentId string, status string) error
}
