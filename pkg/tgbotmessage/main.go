package tgbotmessage

import (
	"context"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type ActionSetterGetter interface {
	SetAction(ctx context.Context, id int64, action Action)
	GetAction(ctx context.Context, id int64) Action
}

type (
	Action      = string
	HandlerFunc = func(upd tgbotapi.Update) (next bool)
)

type Handler struct {
	store ActionSetterGetter

	middleware []HandlerFunc
	handlers   map[Action]*Item

	NotFound HandlerFunc
}

func NewHandler(storage ActionSetterGetter) *Handler {
	return &Handler{
		store:    storage,
		handlers: make(map[Action]*Item),
	}
}

func (h *Handler) Use(mw HandlerFunc) {
	h.middleware = append(h.middleware, mw)
}

func (h *Handler) Add(action Action, handler ...HandlerFunc) {
	_, ok := h.handlers[action]
	if ok {
		h.handlers[action].Handlers = handler
		return
	}

	h.handlers[action] = NewItem(handler, nil)
}

func (h *Handler) On(action Action, on HandlerFunc) {
	_, ok := h.handlers[action]
	if ok {
		h.handlers[action].On = on
		return
	}

	h.handlers[action] = NewItem(nil, on)
}

func (h *Handler) UseNotFound(update tgbotapi.Update) {
	if h.NotFound != nil {
		h.NotFound(update)
	}
}

func (h *Handler) Handle(update tgbotapi.Update) {
	for _, mw := range h.middleware {
		if !mw(update) {
			return
		}
	}

	var action Action

	if update.Message.IsCommand() {
		action = "/" + update.Message.Command()
	} else {
		action = h.GetAction(update.Message.From.ID)
	}

	handlers, ok := h.handlers[action]
	if !ok || handlers.Handlers == nil {
		h.UseNotFound(update)
		return
	}

	for _, handler := range handlers.Handlers {
		if !handler(update) {
			break
		}
	}
}

func (h *Handler) SetAction(id int64, action Action) {
	h.store.SetAction(context.Background(), id, action)
}

func (h *Handler) SetActionWithOn(id int64, action Action, update tgbotapi.Update) error {
	h.SetAction(id, action)

	item, ok := h.handlers[action]
	if !ok || item.On == nil {
		return fmt.Errorf("action %s not found", action)
	}

	item.On(update)
	return nil
}

func (h *Handler) GetAction(id int64) Action {
	return h.store.GetAction(context.Background(), id)
}
