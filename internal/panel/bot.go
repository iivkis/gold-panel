package panel

import (
	"context"
	"gold-panel/config"
	"gold-panel/pkg/tgbotmessage"
	"gold-panel/pkg/tgbotmessage/actstore"

	controller "gold-panel/internal/controller/tg-bot/panel"
	"gold-panel/internal/service/v1"

	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func RunBot(service service.IService) {
	bot, err := tgbotapi.NewBotAPI(config.GetConfig().Panel.BotToken)
	if err != nil {
		panic(err)
	}

	messageHandler := tgbotmessage.NewHandler(actstore.NewStore(
		context.Background(),
		actstore.Options{
			ActionLifetime: time.Hour * 2,
			ActionDefault:  controller.ACT_HOME,
		},
	))

	controller.NewTGBotPanel(bot, service, messageHandler).
		Register()

	u := tgbotapi.NewUpdate(0)
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		user := update.SentFrom()
		if user == nil {
			continue
		}

		if update.Message != nil {
			messageHandler.Handle(update)
		}
	}
}
