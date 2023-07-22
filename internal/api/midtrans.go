package api

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"go-wallet.in/domain"
	"go-wallet.in/dto"
	"go-wallet.in/internal/util"
)

type midtransApi struct {
	midtransService domain.MidtransService
	topupService    domain.TopupService
}

func NewMidtrans(app *fiber.App, midtransService domain.MidtransService, topupService domain.TopupService) {
	api := &midtransApi{
		midtransService: midtransService,
		topupService:    topupService,
	}

	app.Post("/midtrans/payment-callback", api.PaymentHandlerNotification)
}

func (m midtransApi) PaymentHandlerNotification(ctx *fiber.Ctx) error {
	var notificationPayload map[string]interface{}
	if err := ctx.BodyParser(&notificationPayload); err != nil {
		return ctx.Status(400).JSON(util.BuildErrorResponse("failed to parse request body", err))
	}

	orderId, exists := notificationPayload["order_id"].(string)
	if !exists {
		return ctx.Status(400).JSON(util.BuildErrorResponse("failed to parse request body", errors.New("order_id not found")))
	}

	success, _ := m.midtransService.VerifyPayment(ctx.Context(), orderId)
	if success {
		_ = m.topupService.ConfirmedTopup(ctx.Context(), orderId)
		return ctx.Status(util.HttpSatusOk).JSON(util.BuildResponse("success to verify payment", dto.EmptyObj{}))
	}

	return ctx.Status(400).JSON(util.BuildErrorResponse("failed to verify payment", errors.New("payment not verified")))
}
