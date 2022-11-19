package controller

import (
	tgbotcallback "gold-panel/pkg/tgbotcallback"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *controller) OnSendApplication(upd tgbotapi.Update) (next bool) {
	response := tgbotapi.NewMessage(
		upd.SentFrom().ID,
		"Тебя еще нет в списке нашей команде. Желаешь отправить заявку на вступление?",
	)
	response.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Отправить заявку", CB_SEND_APPLICATION_STEP_1),
		),
	)

	c.bot.Send(response)
	return
}

func (c *controller) CallbackApplicationStep1(upd tgbotapi.Update) (next bool) {
	response := tgbotapi.NewMessage(
		upd.SentFrom().ID,
		"Хорошо, я отправлю твою заявку, но сначала тебе нужно заполнить небольшую анкету",
	)
	c.bot.Send(response)

	response = tgbotapi.NewMessage(
		upd.SentFrom().ID,
		"Откуда Вы узнали о нас?",
	)
	response.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"Форум",
				tgbotcallback.NewCallbackData(CB_SEND_APPLICATION_STEP_2, "forum"),
			),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"Пригласили",
				tgbotcallback.NewCallbackData(CB_SEND_APPLICATION_STEP_2, "invited"),
			),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				"Реклама",
				tgbotcallback.NewCallbackData(CB_SEND_APPLICATION_STEP_2, "ad"),
			),
		),
	)

	c.bot.Send(response)
	return
}

func (c *controller) CallbackApplicationStep2(upd tgbotapi.Update) (next bool) {
	_, payload := tgbotcallback.SplitCallback(upd.CallbackData())

	response := tgbotapi.NewMessage(
		upd.SentFrom().ID,
		"",
	)

	switch payload {
	case "forum":
		response.Text = "Введите ссылку на профиль:"
	case "invited":
		response.Text = "Введите ник пригласившего:"
	case "ad":
		response.Text = "Где вы увидели рекламу?"
	}

	c.hanler.SetAction(upd.SentFrom().ID, ACT_SEND_APPLICATION_STEP_2_INPUT)
	c.bot.Send(response)
	return
}

func (c *controller) ApplicationStep2Input(upd tgbotapi.Update) (next bool) {
	return
}
