package repo

import _ "github.com/doug-martin/goqu/v9/dialect/mysql"

func newRepo() IRepo {
	source := "app:app-password@tcp(localhost:3306)/gold_panel"
	return NewRepo(source)
}
