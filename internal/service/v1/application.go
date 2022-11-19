package service

import (
	"context"
	"gold-panel/internal/entity"
	"gold-panel/internal/repo"
	"time"

	"github.com/goccy/go-json"
)

type IApplication interface {
	ApplicationInsert(ctx context.Context, dto *ApplicationInsertDTO) (insertID int64, err error)
	ApplicationGet(ctx context.Context, dto *ApplicationGetDTO) (model *entity.Application, err error)
	ApplicationAccept(ctx context.Context, dto *ApplicationAcceptDTO) (affected int64, err error)
	ApplicationDiscard(ctx context.Context, dto *ApplicationDiscardDTO) (affected int64, err error)
	ApplicationSelect(ctx context.Context, dto *ApplicationSelectDTO) (arr []entity.Application, err error)
	ApplicationFormSave(ctx context.Context, dto *ApplicationFormSaveDTO) error
	ApplicationFormGet(ctx context.Context, dto *ApplicationFormGetDTO) (*entity.ApplicationForm, error)
	ApplicationFormDel(ctx context.Context, dto *ApplicationFormDelDTO) error
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
		&repo.ApplicationGetDTO{
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

type ApplicationFormSaveDTO struct {
	KeyID string
	Form  *entity.ApplicationForm
}

func (s *Service) ApplicationFormSave(ctx context.Context, dto *ApplicationFormSaveDTO) error {
	b, err := json.MarshalContext(ctx, dto.Form)
	if err != nil {
		return err
	}

	status := s.redis.Set(ctx, "application-form-"+dto.KeyID, string(b), time.Hour*24)
	return status.Err()
}

type ApplicationFormGetDTO struct {
	KeyID string
}

func (s *Service) ApplicationFormGet(ctx context.Context, dto *ApplicationFormGetDTO) (*entity.ApplicationForm, error) {
	formstr, err := s.redis.Get(ctx, "application-form-"+dto.KeyID).Result()
	if err != nil {
		return nil, err
	}

	var form entity.ApplicationForm
	if err := json.UnmarshalContext(ctx, []byte(formstr), &form); err != nil {
		return nil, err
	}
	return &form, nil
}

type ApplicationFormDelDTO struct {
	KeyID string
}

func (s *Service) ApplicationFormDel(ctx context.Context, dto *ApplicationFormDelDTO) error {
	err := s.redis.Del(ctx, "application-form-"+dto.KeyID).Err()
	return err
}
