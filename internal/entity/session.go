package entity

type Session struct {
	ID             int64
	Phone          string
	Confirmed      bool
	Is2Fa          bool
	RefUserID      int64
	RefWorkerBotID int64
}
