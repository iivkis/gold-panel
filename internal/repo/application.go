package repo

import (
	"context"
	"gold-panel/internal/entity"
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

func (r *Repo) ApplicationInsert(ctx context.Context, dto *ApplicationInsertDTO) (int64, error) {
	query := `
	INSERT INTO application (
		key_id,
		tag
	) VALUES (?, ?)
	`
	result, err := r.db.ExecContext(ctx, query, dto.KeyID, dto.Tag)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

type ApplicationGetDTO struct {
	KeyID string
}

func (r *Repo) ApplicationGet(ctx context.Context, dto *ApplicationAcceptDTO) (*entity.Application, error) {
	var model entity.Application

	query := `
	SELECT * FROM application
	WHERE key_id = ?
	`

	err := r.db.Get(&model, query, dto.KeyID)
	if err != nil {
		return nil, err
	}

	return &model, nil
}

type ApplicationAcceptDTO struct {
	KeyID string
}

func (r *Repo) ApplicationAccept(ctx context.Context, dto *ApplicationAcceptDTO) (int64, error) {
	query := `
	UPDATE application
	SET
		invited = ?
	WHERE key_id = ?
	`

	result, err := r.db.ExecContext(ctx, query, true, dto.KeyID)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

type ApplicationDiscardDTO struct {
	KeyID string
}

func (r *Repo) ApplicationDiscard(ctx context.Context, dto *ApplicationDiscardDTO) (int64, error) {
	query := `
	UPDATE application
	SET
		invited = ?
	WHERE key_id = ?
	`

	result, err := r.db.ExecContext(ctx, query, false, dto.KeyID)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

type ApplicationSelectDTO struct {
	Limit  uint
	Offset uint
}

func (r *Repo) ApplicationSelect(ctx context.Context, dto *ApplicationSelectDTO) ([]entity.Application, error) {
	var arr []entity.Application

	if dto.Limit == 0 {
		dto.Limit = 1000
	}

	sqlq, _, _ := dialect.
		From("application").
		Limit(dto.Limit).
		Offset(dto.Offset).
		ToSQL()

	err := r.db.Select(&arr, sqlq)
	if err != nil {
		return nil, err
	}

	return arr, err
}
