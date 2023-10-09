package user

import (
	"context"

	"github.com/citcho/go-gizlog/internal/auth/domain/user"
	"github.com/citcho/go-gizlog/internal/user/interface/private/intraprocess"
)

type IntraprocessService struct {
	ic *intraprocess.IntraprocessController
}

func NewIntraprocessService(ic *intraprocess.IntraprocessController) *IntraprocessService {
	return &IntraprocessService{
		ic: ic,
	}
}

func (is IntraprocessService) FetchByEmail(ctx context.Context, email string) (*user.User, error) {
	u, err := is.ic.FetchByEmail(ctx, email)
	if err != nil {
		return &user.User{}, nil
	}

	return authUserFromIntraprocessUser(u)
}

func authUserFromIntraprocessUser(u intraprocess.User) (*user.User, error) {
	return user.NewUser(u.ID, u.Email, u.Password)
}
