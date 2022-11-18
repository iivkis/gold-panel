package entity

import "github.com/goccy/go-json"

type WorkerBotDialogData struct {
}

type WorkerBot struct {
	ID         int64
	Botname    string // bot username (tag)
	DialogData json.RawMessage
	RefUserID  int64
}
