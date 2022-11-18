package workerbot

import (
	"gold-panel/internal/service/v1"
	"gold-panel/pkg/tgbotmessage"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Controller struct {
	bot     *tgbotapi.BotAPI
	handler *tgbotmessage.Handler
	service *service.Service
}

func NewController(bot *tgbotapi.BotAPI, handler *tgbotmessage.Handler, service *service.Service) *Controller {
	return &Controller{
		bot:     bot,
		handler: handler,
		service: service,
	}
}

type HandlerData struct {
	Action string
	Text   string
}

func (c *Controller) Setup(data []HandlerData) {
	for _, d := range data {
		f := func(upd tgbotapi.Update) (next bool) {
			c.bot.Send(
				tgbotapi.NewMessage(
					upd.SentFrom().ID,
					d.Text,
				),
			)
			return
		}
		c.handler.Add(d.Action, f)
	}
}
