package workerbot

import (
	"context"
	"gold-panel/internal/service/v1"
	"gold-panel/pkg/tgbotmessage"
	"gold-panel/pkg/tgbotmessage/actstore"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

type WorkerBot struct {
	workerID int64
	bot      *tgbotapi.BotAPI
	service  service.IService

	ctx    context.Context
	cancel context.CancelFunc
}

func NewWorkerBot(ctx context.Context, workerID int64, service service.IService) (*WorkerBot, error) {
	ctx, cancel := context.WithCancel(ctx)

	w := &WorkerBot{
		workerID: workerID,
		service:  service,

		ctx:    ctx,
		cancel: cancel,
	}

	err := w.init()
	if err != nil {
		return nil, err
	}

	return w, nil
}

func (w *WorkerBot) init() error {
	token, err := w.service.WorkerGetToken(context.Background(), w.workerID)
	if err != nil {
		return err
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return err
	}
	w.bot = bot

	return nil
}

func (w *WorkerBot) Listen() error {
	me, err := w.bot.GetMe()
	if err != nil {
		return err
	}

	logrus.WithFields(logrus.Fields{
		"botname": me.UserName,
	}).Info("Running...")

	message := tgbotmessage.NewHandler(
		actstore.NewStore(
			w.ctx,
			actstore.Options{
				ActionDefault:  "#",
				ActionLifetime: time.Hour,
			},
		),
	)

	u := tgbotapi.NewUpdate(0)
	updates := w.bot.GetUpdatesChan(u)

	go func() {
	listen:
		for {
			select {
			case <-w.ctx.Done():
				break listen

			case update := <-updates:
				user := update.SentFrom()
				if user == nil {
					continue
				}

				logrus.WithFields(logrus.Fields{
					"id":   user.ID,
					"user": user.UserName,
					"bot":  me.UserName,
				}).Info("update")

				if update.Message != nil {
					message.Handle(update)
				}
			}
		}
	}()

	return nil
}

func (w *WorkerBot) Finish() {
	w.cancel()
}
