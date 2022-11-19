package controller

import (
	"gold-panel/internal/service/v1"
	"gold-panel/pkg/tgbotmessage"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type controller struct {
	bot     *tgbotapi.BotAPI
	service service.IService
	hanler  *tgbotmessage.Handler
}

func NewTGBotPanel(bot *tgbotapi.BotAPI, serice service.IService, handler *tgbotmessage.Handler) *controller {
	return &controller{
		bot:     bot,
		service: serice,
		hanler:  handler,
	}
}
