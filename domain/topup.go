package domain

import (
	"context"

	"go-wallet.in/dto"
)

type Topup struct {
	ID      string  `db:"id"`
	UserID  int64   `db:"user_id"`
	Amount  float64 `db:"amount"`
	Status  int8    `db:"status"`
	SnapURL string  `db:"snap_url"`
}

type TopupRepository interface {
	FindByID(ctx context.Context, id string) (Topup, error)
	Insert(ctx context.Context, topup *Topup) error
	Update(ctx context.Context, topup *Topup) error
}

type TopupService interface {
	ConfirmedTopup(ctx context.Context, id string) error
	InitializeTopup(ctx context.Context, req dto.TopupReq) (dto.TopupRes, error)
}
