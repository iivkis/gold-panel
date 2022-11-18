package service

import (
	"context"
	"gold-panel/internal/entity"
	"gold-panel/internal/repo"
)

type IApplication interface {
	ApplicationInsert(ctx context.Context, dto *ApplicationInsertDTO) (insertID int64, err error)
	ApplicationGet(ctx context.Context, dto *ApplicationAcceptDTO) (model *entity.Application, err error)
	ApplicationAccept(ctx context.Context, dto *ApplicationAcceptDTO) (affected int64, err error)
	ApplicationDiscard(ctx context.Context, dto *ApplicationDiscardDTO) (affected int64, err error)
	ApplicationSelect(ctx context.Context, dto *ApplicationSelectDTO) (arr []entity.Application, err error)
}

type ApplicationInsertDTO struct {
	KeyID string
	Tag   string
}

func (s *Service) ApplicationInsert(ctx context.Context, dto *ApplicationInsertDTO) (int64, error) {
	return s.repo.ApplicationInsert(
		ctx,
		&repo.ApplicationInsertDTO{
			KeyID: dto.KeyID,
			Tag:   dto.KeyID,
		},
	)
}

type ApplicationGetDTO struct {
	KeyID string
}

func (s *Service) ApplicationGet(ctx context.Context, dto *ApplicationGetDTO) (*entity.Application, error) {
	return s.repo.ApplicationGet(
		ctx,
		&repo.ApplicationAcceptDTO{
			KeyID: dto.KeyID,
		},
	)
}

type ApplicationAcceptDTO struct {
	KeyID string
}

func (s *Service) ApplicationAccept(ctx context.Context, dto *ApplicationAcceptDTO) (int64, error) {
	return s.repo.ApplicationAccept(
		ctx,
		&repo.ApplicationAcceptDTO{
			KeyID: dto.KeyID,
		},
	)
}

type ApplicationDiscardDTO struct {
	KeyID string
}

func (s *Service) ApplicationDiscard(ctx context.Context, dto *ApplicationDiscardDTO) (int64, error) {
	return s.repo.ApplicationDiscard(
		ctx,
		&repo.ApplicationDiscardDTO{
			KeyID: dto.KeyID,
		},
	)
}

type ApplicationSelectDTO struct {
	Limit  uint
	Offset uint
}

func (s *Service) ApplicationSelect(ctx context.Context, dto *ApplicationSelectDTO) ([]entity.Application, error) {
	return s.repo.ApplicationSelect(
		ctx,
		&repo.ApplicationSelectDTO{
			Limit:  dto.Limit,
			Offset: dto.Offset,
		},
	)
}
