package application

import (
	"context"
	"fmt"

	"github.com/citcho/go-gizlog/internal/auth/domain/user"
	"github.com/citcho/go-gizlog/internal/common/derror"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	FetchByEmail(context.Context, string) (user.User, error)
}

type TokenGenerator interface {
	GenerateToken(ctx context.Context, u user.User) ([]byte, error)
}

type IAuthUsecase interface {
	Login(context.Context, LoginCommand) (string, error)
}

type AuthUsecase struct {
	tg TokenGenerator
	us IUserService
}

func NewAuthUsecase(us IUserService, tg TokenGenerator) *AuthUsecase {
	return &AuthUsecase{
		tg: tg,
		us: us,
	}
}

func (au *AuthUsecase) Login(ctx context.Context, cmd LoginCommand) (token string, err error) {
	defer derror.Wrap(&err, "Login(%+v)", cmd)

	u, err := au.us.FetchByEmail(ctx, cmd.Email)
	if err != nil {
		return "", err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(u.Password()), []byte(cmd.Password)); err != nil {
		return "", err
	}

	t, err := au.tg.GenerateToken(ctx, u)
	if err != nil {
		return "", fmt.Errorf("failed to generate JWT: %w", err)
	}

	token = string(t)

	return token, nil
}
