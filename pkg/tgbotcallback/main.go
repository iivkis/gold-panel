package tgbotcallback

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type (
	HandlerFunc = func(upd tgbotapi.Update)
)

type Handler struct {
	handlers map[string]HandlerFunc
}

func NewHandler() *Handler {
	return &Handler{
		handlers: make(map[string]HandlerFunc),
	}
}

func (h *Handler) Add(callback string, handler HandlerFunc) {
	h.handlers[callback] = handler
}

func (h *Handler) Handle(upd tgbotapi.Update) error {
	callback, _ := SplitCallback(upd.CallbackData())

	if handler, ok := h.handlers[callback]; ok {
		handler(upd)
		return nil
	}

	return fmt.Errorf("undefined callback: %s", callback)
}
