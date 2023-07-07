package sse

import (
	"bufio"
	"encoding/json"
	"fmt"
	"go-bank/dto"

	"github.com/gofiber/fiber/v2"
)

type notificationSSE struct {
	hub *dto.Hub
}

func NewNotification(app *fiber.App, authMid fiber.Handler, hub *dto.Hub) {
	api := &notificationSSE{
		hub: hub,
	}

	app.Get("sse/notification-stream", authMid, api.StreamNotification)
}

func (n notificationSSE) StreamNotification(ctx *fiber.Ctx) error {
	ctx.Set("Content-Type", "text/event-stream")

	user := ctx.Locals("x-user").(*dto.UserData)
	n.hub.NotificationChannel[user.ID] = make(chan dto.NotificationData)

	ctx.Context().SetBodyStreamWriter(func(w *bufio.Writer) {
		event := fmt.Sprint("event: %s/n"+"data: /n/n", "initial")
		_, _ = fmt.Fprint(w, event)
		_ = w.Flush()
		for notification := range n.hub.NotificationChannel[user.ID] {
			data, _ := json.Marshal(notification)

			event := fmt.Sprint("event: %s/n"+"data: %s/n/n", "notification-updated", data)
			_, _ = fmt.Fprint(w, event)
			_ = w.Flush()
		}
	})

	return nil
}
