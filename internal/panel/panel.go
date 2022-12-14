package panel

import (
	"gold-panel/config"
	"gold-panel/internal/repo"
	service "gold-panel/internal/service/v1"

	"github.com/go-redis/redis/v9"
)

func Launch() {
	redisdb := redis.NewClient(&redis.Options{
		Addr: "redisdb:6379",
	})

	repository := repo.NewRepo(
		&repo.Source{
			Host:     "mysqldb",
			Port:     "3306",
			User:     config.GetEnv().MySQLUser,
			Password: config.GetEnv().MySQLPassword,
			DBName:   config.GetEnv().MySQLDBName,
		},
	)

	service := service.NewService(repository, redisdb)
	RunBot(service)
}
