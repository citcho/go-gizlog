package intraprocess

import (
	"context"

	"github.com/citcho/go-gizlog/internal/user/domain/user"
)

type IUserUsecase interface {
	FetchByEmail(context.Context, string) (*user.User, error)
}

type UserController struct {
	usecase IUserUsecase
}

func NewUserController(uu IUserUsecase) *UserController {
	return &UserController{
		usecase: uu,
	}
}

func (uc UserController) FetchByEmail(ctx context.Context, email string) (*user.User, error) {
	u, err := uc.usecase.FetchByEmail(ctx, email)
	if err != nil {
		return &user.User{}, err
	}

	return u, nil
}
