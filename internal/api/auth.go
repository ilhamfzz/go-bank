package api

import (
	"go-wallet.in/domain"
	"go-wallet.in/dto"
	"go-wallet.in/internal/util"

	"github.com/gofiber/fiber/v2"
)

type authApi struct {
	userService domain.UserService
}

func NewAuth(app *fiber.App, userService domain.UserService, authMid fiber.Handler) {
	api := authApi{
		userService: userService,
	}

	app.Post("user/register", api.Register)
	app.Post("user/validate-otp", api.ValidateOTP)
	app.Post("token/generate", api.GenerateToken)
	app.Get("token/validate", authMid, api.ValidateToken)
}

func (a authApi) Register(ctx *fiber.Ctx) error {
	var req dto.RegisterReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(400).JSON(util.BuildErrorResponse("failed to parse request body", err))
	}

	user, err := a.userService.Register(ctx.Context(), req)
	if err != nil {
		return ctx.Status(util.GetErrHttpStatusCode(err)).JSON(util.BuildErrorResponse("failed to register", err))
	}

	return ctx.Status(util.HttpSatusCreated).JSON(util.BuildResponse("success to register user", user))
}

func (a authApi) GenerateToken(ctx *fiber.Ctx) error {
	var req dto.AuthReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(400).JSON(util.BuildErrorResponse("failed to parse request body", err))
	}

	token, err := a.userService.Authenticate(ctx.Context(), req)
	if err != nil {
		return ctx.Status(util.GetErrHttpStatusCode(err)).JSON(util.BuildErrorResponse("failed to authenticate", err))
	}

	return ctx.Status(util.HttpSatusOk).JSON(util.BuildResponse("success to generate token", token))
}

func (a authApi) ValidateToken(ctx *fiber.Ctx) error {
	user := ctx.Locals("x-user")
	return ctx.Status(util.HttpSatusOk).JSON(util.BuildResponse("success to validate token", user))
}

func (a authApi) ValidateOTP(ctx *fiber.Ctx) error {
	var req dto.ValidateOTPReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(400).JSON(util.BuildErrorResponse("failed to parse request body", err))
	}

	err := a.userService.ValidateOTP(ctx.Context(), req)
	if err != nil {
		return ctx.Status(util.GetErrHttpStatusCode(err)).JSON(util.BuildErrorResponse("failed to validate otp", err))
	}

	return ctx.Status(util.HttpSatusOk).JSON(util.BuildResponse("success to validate otp", "Your account has been activated"))
}
