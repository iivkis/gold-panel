package tgbotmessage

type Item struct {
	On       HandlerFunc
	Handlers []HandlerFunc
}

func NewItem(handlers []HandlerFunc, on HandlerFunc) *Item {
	return &Item{
		On:       on,
		Handlers: handlers,
	}
}
