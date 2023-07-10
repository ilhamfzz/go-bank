package repository

import (
	"context"
	"database/sql"

	"go-bank/domain"

	"github.com/doug-martin/goqu/v9"
)

type AccountRepository struct {
	db *goqu.Database
}

func NewAccount(con *sql.DB) domain.AccountRepository {
	return &AccountRepository{
		db: goqu.New("default", con),
	}
}

func (a *AccountRepository) FindByUserID(ctx context.Context, id int64) (account domain.Account, err error) {
	dataset := a.db.From("accounts").Where(goqu.Ex{
		"user_id": id,
	})

	_, err = dataset.ScanStructContext(ctx, &account)
	return
}

func (a *AccountRepository) FindByAccountNumber(ctx context.Context, accountNumber string) (account domain.Account, err error) {
	dataset := a.db.From("accounts").Where(goqu.Ex{
		"account_number": accountNumber,
	})

	_, err = dataset.ScanStructContext(ctx, &account)
	return
}

func (a *AccountRepository) UpdateAccountBalance(ctx context.Context, account *domain.Account) error {
	dataset := a.db.Update("accounts").Where(goqu.Ex{
		"id": account.ID,
	}).Set(goqu.Record{
		"balance": account.Balance,
	}).Executor()

	_, err := dataset.ExecContext(ctx)
	return err
}

func (a *AccountRepository) Insert(ctx context.Context, account *domain.Account) (err error) {
	dataset := a.db.Insert("accounts").Rows(
		goqu.Record{
			"user_id":        account.UserId,
			"account_number": account.AccountNumber,
			"balance":        account.Balance,
		}).Executor()

	_, err = dataset.ExecContext(ctx)
	return
}
