package user

import (
	"context"

	"github.com/citcho/go-gizlog/internal/user/domain/user"
	"github.com/citcho/go-gizlog/internal/user/infrastructure/dao"
	"github.com/uptrace/bun"
)

type MySQLRepository struct {
	db *bun.DB
}

func NewMySQLRepository(db *bun.DB) *MySQLRepository {
	return &MySQLRepository{
		db: db,
	}
}

func (mr *MySQLRepository) Save(ctx context.Context, user *user.User) error {
	u := dao.User{
		ID:       user.ID(),
		Name:     user.Name(),
		Email:    user.Email(),
		Password: user.Password(),
	}

	_, err := mr.db.NewInsert().
		Model(&u).
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (mr *MySQLRepository) Exists(ctx context.Context, user *user.User) (bool, error) {
	exists, err := mr.db.NewSelect().
		Model((*dao.User)(nil)).
		Where("email = ?", user.Email()).
		Exists(ctx)
	if err != nil {
		return exists, err
	}

	return exists, nil
}

func (mr *MySQLRepository) FetchByEmail(ctx context.Context, email string) (*dao.User, error) {
	var u dao.User
	err := mr.db.NewSelect().
		Model(&u).
		Where("email = ?", email).
		Scan(ctx)
	if err != nil {
		return &dao.User{}, err
	}

	return &u, nil
}
