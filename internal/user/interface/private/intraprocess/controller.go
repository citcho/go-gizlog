package intraprocess

import (
	"context"

	"github.com/citcho/go-gizlog/internal/user/domain/user"
)

type IUserRepository interface {
	FetchByEmail(context.Context, string) (*user.User, error)
}

type User struct {
	ID       string
	Name     string
	Email    string
	Password string
}

func UserFromDomainUser(domainUser *user.User) User {
	return User{
		ID:       domainUser.ID(),
		Name:     domainUser.Name(),
		Email:    domainUser.Email(),
		Password: domainUser.Password(),
	}
}

type IUserUsecase interface {
	FetchByEmail(context.Context, string) (*user.User, error)
}

type IntraprocessController struct {
	repository IUserRepository
}

func NewIntraprocessController(ur IUserRepository) *IntraprocessController {
	return &IntraprocessController{
		repository: ur,
	}
}

func (uc IntraprocessController) FetchByEmail(ctx context.Context, email string) (User, error) {
	domainUser, err := uc.repository.FetchByEmail(ctx, email)
	if err != nil {
		return User{}, err
	}

	return UserFromDomainUser(domainUser), nil
}
