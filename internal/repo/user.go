package repo

import "context"

type IUserRepo interface {
	UserCreate(ctx context.Context) error
}

type UserCreateDTO struct {
	Username string
}

func (r *Repo) UserCreate(ctx context.Context) error {
	return nil
}
