package api

import (
	"go-bank/domain"
	"go-bank/dto"
	"go-bank/internal/util"

	"github.com/gofiber/fiber/v2"
)

type transferApi struct {
	transferService domain.TransactionService
}

func NewTransfer(app *fiber.App, authMid fiber.Handler, transactionService domain.TransactionService) {
	api := &transferApi{
		transferService: transactionService,
	}

	transfer := app.Group("/transfer", authMid)
	transfer.Post("/inquiry", api.TransferInquiry)
	transfer.Post("/execute", api.TransferExecute)
}

func (t transferApi) TransferInquiry(ctx *fiber.Ctx) error {
	var req dto.TransferInquiryReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(400).JSON(util.BuildErrorResponse("failed to parse request body", err))
	}

	inquerry, err := t.transferService.TransferInquiry(ctx.Context(), req)
	if err != nil {
		return ctx.Status(util.GetErrHttpStatusCode(err)).JSON(util.BuildErrorResponse("failed to get transfer inquiry", err))
	}

	return ctx.Status(util.HttpSatusOk).JSON(util.BuildResponse("success to get transfer inquiry", inquerry))
}

func (t transferApi) TransferExecute(ctx *fiber.Ctx) error {
	var req dto.TransferExecuteReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(400).JSON(util.BuildErrorResponse("failed to parse request body", err))
	}

	err := t.transferService.TransferExecute(ctx.Context(), req)
	if err != nil {
		return ctx.Status(util.GetErrHttpStatusCode(err)).JSON(util.BuildErrorResponse("failed to execute transfer", err))
	}

	return ctx.Status(util.HttpSatusOk).JSON(util.BuildResponse("success to execute transfer", "Your transfer has been successfully executed"))
}
