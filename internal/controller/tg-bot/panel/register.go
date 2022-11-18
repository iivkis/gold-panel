package controller

import "gold-panel/pkg/tgbotmessage"

func (c *controller) Register(h *tgbotmessage.Handler) {
	h.Add("/me", c.Commad.Me)
}
