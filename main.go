package main

import (
	"go-wallet.in/dto"
	"go-wallet.in/internal/api"
	"go-wallet.in/internal/component"
	"go-wallet.in/internal/config"
	"go-wallet.in/internal/middleware"
	"go-wallet.in/internal/repository"
	"go-wallet.in/internal/service"

	"go-wallet.in/internal/sse"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cnf := config.Get()
	dbConnection := component.DatabaseConnection(cnf)
	// cacheConnection := component.GetCacheConnection()
	cacheConnection := repository.NewRedisClient(cnf)

	hub := &dto.Hub{
		NotificationChannel: map[int64]chan dto.NotificationData{},
	}

	userRepository := repository.NewUser(dbConnection)
	accountRepository := repository.NewAccount(dbConnection)
	transactionRepository := repository.NewTransaction(dbConnection)
	notificationRepository := repository.NewNotification(dbConnection)

	emailService := service.NewEmail(cnf)
	userService := service.NewUser(userRepository, cacheConnection, emailService)
	transactionService := service.NewTransaction(accountRepository, transactionRepository, cacheConnection, notificationRepository, hub)
	notificationService := service.NewNotification(notificationRepository)

	authMiddleware := middleware.Authenticate(userService)

	app := fiber.New()
	api.NewAuth(app, userService, authMiddleware)
	api.NewTransfer(app, authMiddleware, transactionService)
	api.NewNotification(app, authMiddleware, notificationService)

	sse.NewNotification(app, authMiddleware, hub)

	_ = app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
}
