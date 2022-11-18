package service

import "gold-panel/internal/repo"

//go:generate mockgen -destination=./mock/service.go -package=mock_service gold-panel/internal/service/v1 IService
type IService interface {
	IWorkerService
}

type Service struct {
	repo repo.IRepo
}

func NewService() IService {
	return &Service{}
}
