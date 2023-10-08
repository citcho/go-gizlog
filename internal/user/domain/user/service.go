package user

import (
	"context"
)

type IUserRepository interface {
	Exists(context.Context, *User) (bool, error)
}

type UserService struct {
	repository IUserRepository
}

func NewUserService(ur IUserRepository) *UserService {
	return &UserService{
		repository: ur,
	}
}

func (us UserService) Exists(ctx context.Context, r *User) (bool, error) {
	exists, err := us.repository.Exists(ctx, r)
	if err != nil {
		return exists, err
	}

	return exists, nil
}
