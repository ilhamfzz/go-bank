package service

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"go-wallet.in/domain"
	"go-wallet.in/dto"
)

type topupService struct {
	midtransService  domain.MidtransService
	topupRepo        domain.TopupRepository
	accountRepo      domain.AccountRepository
	notificationRepo domain.NotificationRepository
	transactionRepo  domain.TransactionRepository
}

func NewTopup(midtransService domain.MidtransService, topupRepo domain.TopupRepository, accountRepo domain.AccountRepository, notificationRepo domain.NotificationRepository, transactionRepo domain.TransactionRepository) domain.TopupService {
	return &topupService{
		midtransService:  midtransService,
		topupRepo:        topupRepo,
		accountRepo:      accountRepo,
		notificationRepo: notificationRepo,
		transactionRepo:  transactionRepo,
	}
}

func (t *topupService) InitializeTopup(ctx context.Context, req dto.TopupReq) (dto.TopupRes, error) {
	topup := domain.Topup{
		ID:     uuid.NewString(),
		UserID: req.UserID,
		Status: 0,
		Amount: req.Amount,
	}

	err := t.midtransService.GenerateSnapURL(ctx, &topup)
	if err != nil {
		return dto.TopupRes{}, err
	}

	err = t.topupRepo.Insert(ctx, &topup)
	if err != nil {
		return dto.TopupRes{}, err
	}

	return dto.TopupRes{
		SnapURL: topup.SnapURL,
	}, nil
}

func (t *topupService) ConfirmedTopup(ctx context.Context, id string) error {
	topup, err := t.topupRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if topup == (domain.Topup{}) {
		return domain.ErrTopupNotFound
	}

	account, err := t.accountRepo.FindByUserID(ctx, topup.UserID)
	if err != nil {
		return err
	}
	if account == (domain.Account{}) {
		return domain.ErrAccountNotFound
	}

	err = t.transactionRepo.Insert(ctx, &domain.Transaction{
		AccountId:       account.ID,
		SoftNumber:      "00",
		DofNumber:       account.AccountNumber,
		TransactionType: "C",
		Amount:          topup.Amount,
		TransactionDate: time.Now(),
	})
	if err != nil {
		return err
	}

	account.Balance += topup.Amount
	err = t.accountRepo.UpdateAccountBalance(ctx, &account)
	if err != nil {
		return err
	}

	_ = t.notificationRepo.Insert(ctx, &domain.Notification{
		UserID:    topup.UserID,
		Title:     "Topup Berhasil Diterima",
		Body:      fmt.Sprintf("Topup anda senilai %.2f berhasil diterima", topup.Amount),
		IsRead:    0,
		Status:    1,
		CreatedAt: time.Now(),
	})

	return err
}
