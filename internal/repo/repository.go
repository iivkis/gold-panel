package repo

import (
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type IRepo interface {
	IUserRepo
	IApplication
}

type Repo struct {
	db *sqlx.DB
}

var dialect = goqu.Dialect("mysql")

func NewRepo(source string) IRepo {
	db, err := sqlx.Connect("mysql", source)
	if err != nil {
		panic(err)
	}
	return &Repo{db: db}
}
