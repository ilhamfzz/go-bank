package repository

import (
	"context"
	"database/sql"
	"go-bank/domain"

	"github.com/doug-martin/goqu/v9"
)

type notificationRepository struct {
	db *goqu.Database
}

func NewNotification(con *sql.DB) domain.NotificationRepository {
	return &notificationRepository{
		db: goqu.New("default", con),
	}
}

func (n *notificationRepository) FindByUserID(ctx context.Context, userID int64) (notifications []domain.Notification, err error) {
	dataset := n.db.From("notifications").Where(goqu.Ex{
		"user_id": userID,
	}).Order(goqu.I("created_at").Desc()).Limit(15)

	err = dataset.ScanStructsContext(ctx, &notifications)
	return
}

func (n *notificationRepository) Insert(ctx context.Context, notification *domain.Notification) error {
	dataset := n.db.Insert("notifications").Rows(goqu.Record{
		"user_id":    notification.UserID,
		"status":     notification.Status,
		"title":      notification.Title,
		"body":       notification.Body,
		"is_read":    notification.IsRead,
		"created_at": notification.CreatedAt,
	}).Returning("id").Executor()

	_, err := dataset.ScanStructContext(ctx, notification)
	return err
}

func (n *notificationRepository) Update(ctx context.Context, notification *domain.Notification) error {
	dataset := n.db.Update("notifications").Where(goqu.Ex{
		"id": notification.ID,
	}).Set(goqu.Record{
		"title":   notification.Title,
		"body":    notification.Body,
		"status":  notification.Status,
		"is_read": notification.IsRead,
	}).Executor()

	_, err := dataset.ExecContext(ctx)
	return err
}
