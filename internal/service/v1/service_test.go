package service

import (
	"gold-panel/config"
	"gold-panel/internal/repo"
	"testing"

	"github.com/go-redis/redis/v9"
)

func init() {
	config.LoadConfigFrom("../../../config")
	config.LoadEnvFrom("../../../")
}

func newService(t *testing.T) *Service {
	redisdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	repository := repo.NewRepo(
		&repo.Source{
			Host:     "localhost",
			Port:     "3306",
			User:     config.GetEnv().MySQLUser,
			Password: config.GetEnv().MySQLPassword,
			DBName:   config.GetEnv().MySQLDBName,
		},
	)

	return &Service{
		repo:  repository,
		redis: redisdb,
	}
}

func TestNewService(t *testing.T) {
	newService(t)
}
