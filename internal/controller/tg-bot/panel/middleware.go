package controller

import (
	"context"
	"gold-panel/internal/service/v1"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *controller) MiddlewareApplication(upd tgbotapi.Update) (next bool) {
	application, _ := c.service.ApplicationGet(
		context.Background(),
		&service.ApplicationGetDTO{
			KeyID: "tg-" + strconv.FormatInt(upd.SentFrom().ID, 10),
		},
	)

	if application.Invited {
		return true
	}

	return
}
