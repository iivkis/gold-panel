package repo

import _ "github.com/doug-martin/goqu/v9/dialect/mysql"

func newRepo() IRepo {
	return NewRepo(&Source{
		User:     "app",
		Password: "app-password",
		Host:     "localhost",
		Port:     "3306",
		DBName:   "gold_panel",
	})
}
