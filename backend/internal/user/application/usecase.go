package application

import (
	"context"
	"fmt"

	"github.com/citcho/go-gizlog/internal/common/derror"
	"github.com/citcho/go-gizlog/internal/user/domain/user"
)

type IUserUsecase interface {
	StoreUser(context.Context, StoreUserCommand) error
	FetchByEmail(context.Context, string) (*user.User, error)
}

type UserUsecase struct {
	repository user.IUserRepository
}

func NewUserUsecase(ur user.IUserRepository) *UserUsecase {
	return &UserUsecase{
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
		// TODO: ドメインの知識が漏れ出ているが致し方なしか…？
		return fmt.Errorf("%sは既に登録されています。", u.Email())
	} else {
		if err := uu.repository.Save(ctx, u); err != nil {
			return err
		}
	}

	return nil
}

func (uc UserUsecase) FetchByEmail(ctx context.Context, email string) (*user.User, error) {
	dao, err := uc.repository.FetchByEmail(ctx, email)
	if err != nil {
		return &user.User{}, err
	}

	u, err := user.ReConstruct(
		dao.ID,
		dao.Name,
		dao.Email,
		dao.Password,
	)
	if err != nil {
		return &user.User{}, err
	}

	return u, nil
}
