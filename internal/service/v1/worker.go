package service

import "context"

type IWorkerService interface {
	WorkerAdd(ctx context.Context, dto *WorkerAddDTO) error
	WorkerGetToken(ctx context.Context, workerID int64) (token string, err error)
	WorkerGetAllID(ctx context.Context) ([]int64, error)
}

type WorkerAddDTO struct {
	ID       int64
	Token    string
	Username string
}

func (s *Service) WorkerAdd(ctx context.Context, dto *WorkerAddDTO) error {
	return nil
}

func (s *Service) WorkerGetToken(ctx context.Context, workerID int64) (token string, err error) {
	return "", nil
}

func (s *Service) WorkerGetAllID(ctx context.Context) ([]int64, error) {
	return []int64{}, nil
}
