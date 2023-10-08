package user

import (
	"context"
	"errors"
)

var ErrNotFound = errors.New("user not found")

type IUserRepository interface {
	Save(context.Context, *User) error
	Exists(context.Context, *User) (bool, error)
}
