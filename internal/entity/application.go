package entity

//заявка на доступ к ресурсу
type Application struct {
	ID      int64  `db:"id"`
	KeyID   string `db:"key_id"` // example: "telegram-[user_id]"
	Tag     string `db:"tag"`    // tag or username
	Invited bool   `db:"invited"`
}
