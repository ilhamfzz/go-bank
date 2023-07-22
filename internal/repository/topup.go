package repository

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	"go-wallet.in/domain"
)

type topupRepository struct {
	db *goqu.Database
}

func NewTopup(conn *sql.DB) domain.TopupRepository {
	return &topupRepository{
		db: goqu.New("default", conn),
	}
}

func (t *topupRepository) FindByID(ctx context.Context, id string) (topup domain.Topup, err error) {
	dataset := t.db.From("topups").Where(goqu.Ex{
		"id": id,
	})

	_, err = dataset.ScanStructContext(ctx, &topup)
	return
}

func (t *topupRepository) Insert(ctx context.Context, topup *domain.Topup) error {
	dataset := t.db.Insert("topups").Rows(goqu.Record{
		"id":       topup.ID,
		"user_id":  topup.UserID,
		"amount":   topup.Amount,
		"status":   topup.Status,
		"snap_url": topup.SnapURL,
	}).Executor()

	_, err := dataset.ExecContext(ctx)
	return err
}

func (t *topupRepository) Update(ctx context.Context, topup *domain.Topup) error {
	dataset := t.db.Update("topups").Where(goqu.Ex{
		"id": topup.ID,
	}).Set(goqu.Record{
		"amount":   topup.Amount,
		"status":   topup.Status,
		"snap_url": topup.SnapURL,
	}).Executor()

	_, err := dataset.ExecContext(ctx)
	return err
}
