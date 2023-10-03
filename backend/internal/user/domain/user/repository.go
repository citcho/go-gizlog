package user

import (
	"context"
	"errors"

	"github.com/citcho/go-gizlog/internal/user/infrastructure/dao"
)

var ErrNotFound = errors.New("user not found")

type IUserRepository interface {
	Save(context.Context, *User) error
	Exists(context.Context, *User) (bool, error)
	FetchByEmail(context.Context, string) (*dao.User, error)
}
