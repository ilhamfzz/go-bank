package service

import (
	"context"

	"go-wallet.in/domain"
	"go-wallet.in/dto"
)

type notificationService struct {
	notificationRepo domain.NotificationRepository
}

func NewNotification(notificationRepo domain.NotificationRepository) domain.NotificationService {
	return &notificationService{
		notificationRepo: notificationRepo,
	}
}

func (n notificationService) FindByUserID(ctx context.Context, userID int64) ([]dto.NotificationData, error) {
	notifications, err := n.notificationRepo.FindByUserID(ctx, userID)
	if err != nil {
		return []dto.NotificationData{}, err
	}

	var notificationData []dto.NotificationData
	for _, notification := range notifications {
		notificationData = append(notificationData, dto.NotificationData{
			ID:        notification.ID,
			Title:     notification.Title,
			Body:      notification.Body,
			Status:    notification.Status,
			IsRead:    notification.IsRead,
			CreatedAt: notification.CreatedAt,
		})
	}

	if notificationData == nil {
		notificationData = make([]dto.NotificationData, 0)
	}

	return notificationData, nil
}
