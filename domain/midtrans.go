package domain

import "context"

type MidtransService interface {
	GenerateSnapURL(ctx context.Context, topup *Topup) error
	VerifyPayment(ctx context.Context, orderId string) (bool, error)
}
