package api

import (
	"github.com/gofiber/fiber/v2"
	"go-wallet.in/domain"
	"go-wallet.in/dto"

	"go-wallet.in/internal/util"
)

type topupApi struct {
	topupService domain.TopupService
}

func NewTopup(app *fiber.App, authMid fiber.Handler, topupService domain.TopupService) {
	api := &topupApi{
		topupService: topupService,
	}

	app.Post("/topup/initialize", authMid, api.InitializeTopup)
}

func (t topupApi) InitializeTopup(ctx *fiber.Ctx) error {
	var req dto.TopupReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(400).JSON(util.BuildErrorResponse("failed to parse request body", err))
	}

	user := ctx.Locals("x-user").(dto.UserData)
	req.UserID = user.ID

	topup, err := t.topupService.InitializeTopup(ctx.Context(), req)
	if err != nil {
		return ctx.Status(util.GetErrHttpStatusCode(err)).JSON(util.BuildErrorResponse("failed to initialize topup", err))
	}

	return ctx.Status(util.HttpSatusOk).JSON(util.BuildResponse("success to initialize topup", topup))
}
