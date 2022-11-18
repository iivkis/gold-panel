package controller

import (
	"gold-panel/internal/service/v1"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type controller struct {
	Commad *CommandController
}

func NewTGBotPanel(bot *tgbotapi.BotAPI, serice service.IService) *controller {
	return &controller{
		Commad: NewCommandController(bot),
	}
}
