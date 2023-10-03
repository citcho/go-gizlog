package user

import (
	"context"

	"github.com/citcho/go-gizlog/internal/auth/domain/user"
	"github.com/citcho/go-gizlog/internal/user/interface/private/intraprocess"
)

type UserService struct {
	uc *intraprocess.UserController
}

func NewUserService(uc *intraprocess.UserController) *UserService {
	return &UserService{uc}
}

func (us UserService) FetchByEmail(ctx context.Context, email string) (user.User, error) {
	u, err := us.uc.FetchByEmail(ctx, email)
	if err != nil {
		return user.User{}, nil
	}

	authUser, err := user.NewUser(
		u.ID(),
		u.Email(),
		u.Password(),
	)
	if err != nil {
		return user.User{}, err
	}

	return *authUser, nil
}
