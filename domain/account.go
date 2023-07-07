package domain

import "context"

type Account struct {
	ID            int64   `db:"id"`
	UserId        int64   `db:"user_id"`
	AccountNumber string  `db:"account_number"`
	Balance       float64 `db:"balance"`
}

type AccountRepository interface {
	FindByUserID(ctx context.Context, id int64) (Account, error)
	FindByAccountNumber(ctx context.Context, accountNumber string) (Account, error)
	UpdateAccountBalance(ctx context.Context, id int64, balance float64) error
	Insert(ctx context.Context, account Account) error
}
