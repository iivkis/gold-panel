package repo

import (
	"fmt"

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

type Source struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
}

func NewRepo(source *Source) IRepo {
	src := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", source.User, source.Password, source.Host, source.Port, source.DBName)

	fmt.Println(src)

	db, err := sqlx.Connect("mysql", src)
	if err != nil {
		panic(err)
	}
	return &Repo{db: db}
}
