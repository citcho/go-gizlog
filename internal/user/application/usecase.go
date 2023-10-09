package application

import (
	"context"
	"fmt"

	"github.com/citcho/go-gizlog/internal/common/derror"
	"github.com/citcho/go-gizlog/internal/user/domain/user"
)

type IUserRepository interface {
	Save(context.Context, *user.User) error
	Exists(context.Context, *user.User) (bool, error)
	FetchByEmail(context.Context, string) (*user.User, error)
}

type IUserService interface {
	Exists(context.Context, *user.User) (bool, error)
}

type UserUsecase struct {
	service    IUserService
	repository IUserRepository
}

func NewUserUsecase(us IUserService, ur IUserRepository) *UserUsecase {
	return &UserUsecase{
		service:    us,
		repository: ur,
	}
}

func (uu UserUsecase) StoreUser(ctx context.Context, cmd StoreUserCommand) (err error) {
	defer derror.Wrap(&err, "StoreUser(%+v)", cmd)

	u, err := user.NewUser(
		cmd.ID,
		cmd.Name,
		cmd.Email,
		cmd.Password,
	)
	if err != nil {
		return fmt.Errorf("%s: %w", err, derror.InvalidArgument)
	}

	exists, err := uu.repository.Exists(ctx, u)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("既にこのメールアドレスを持つユーザーが存在しています。: %w", derror.InvalidArgument)
	}

	if err := uu.repository.Save(ctx, u); err != nil {
		return err
	}

	return nil
}

func (uc UserUsecase) FetchByEmail(ctx context.Context, email string) (*user.User, error) {
	u, err := uc.repository.FetchByEmail(ctx, email)
	if err != nil {
		return &user.User{}, err
	}

	return u, nil
}
