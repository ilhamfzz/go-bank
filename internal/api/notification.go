package api

import (
	"context"
	"go-bank/domain"
	"go-bank/dto"
	"go-bank/internal/util"
	"time"

	"github.com/gofiber/fiber/v2"
)

type NotificationApi struct {
	notificationService domain.NotificationService
}

func NewNotification(app *fiber.App, authMid fiber.Handler, notificationService domain.NotificationService) {
	api := &NotificationApi{
		notificationService: notificationService,
	}

	app.Get("/notifications", authMid, api.GetUserNotification)
}

func (api *NotificationApi) GetUserNotification(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 15*time.Second)
	defer cancel()

	user := ctx.Locals("x-user").(dto.UserData)

	notifications, err := api.notificationService.FindByUserID(c, user.ID)
	if err != nil {
		return ctx.Status(util.GetErrHttpStatusCode(err)).JSON(util.BuildErrorResponse("Failed to get user notification", err))
	}

	return ctx.Status(util.HttpSatusOk).JSON(util.BuildResponse("User notification retrieved successfully", notifications))
}
