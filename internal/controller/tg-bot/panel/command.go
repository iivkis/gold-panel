package controller

import (
	"fmt"
	"gold-panel/pkg/tgbotmessage"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type CommandController struct {
	bot *tgbotapi.BotAPI
}

func NewCommandController(bot *tgbotapi.BotAPI) *CommandController {
	return &CommandController{
		bot: bot,
	}
}

func (c *CommandController) Me(upd tgbotapi.Update) (next bool) {
	resp := tgbotapi.NewMessage(
		upd.SentFrom().ID,
		fmt.Sprintf("ID: <code>%d</code>", upd.SentFrom().ID),
	)
	resp.ParseMode = tgbotmessage.PARSE_MODE_HTML

	c.bot.Send(resp)
	return
}
