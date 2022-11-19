package service

import (
	"gold-panel/internal/repo"

	"github.com/go-redis/redis/v9"
)

//go:generate mockgen -destination=./mock/service.go -package=mock_service gold-panel/internal/service/v1 IService
type IService interface {
	IApplication
	IWorkerService
}

type Service struct {
	repo  repo.IRepo
	redis *redis.Client
}

func NewService(repo repo.IRepo, redisdb *redis.Client) IService {
	return &Service{
		repo:  repo,
		redis: redisdb,
	}
}
