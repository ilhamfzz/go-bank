package domain

import (
	"context"
	"time"

	"go-wallet.in/dto"
)

type Transaction struct {
	ID              int64     `db:"id"`
	AccountId       int64     `db:"account_id"`
	SoftNumber      string    `db:"soft_number"`
	DofNumber       string    `db:"dof_number"`
	TransactionType string    `db:"transaction_type"`
	Amount          float64   `db:"amount"`
	TransactionDate time.Time `db:"transaction_date"`
}

type TransactionRepository interface {
	Insert(ctx context.Context, transaction *Transaction) error
}

type TransactionService interface {
	TransferInquiry(ctx context.Context, req dto.TransferInquiryReq) (dto.TransferInquiryRes, error)
	TransferExecute(ctx context.Context, req dto.TransferExecuteReq) error
}
