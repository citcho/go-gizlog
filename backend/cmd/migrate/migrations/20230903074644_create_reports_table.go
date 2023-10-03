package migrations

import (
	"context"

	"github.com/citcho/go-gizlog/internal/report/infrastructure/dao"
	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		_, err := db.NewCreateTable().
			Model((*dao.Report)(nil)).
			Exec(ctx)
		if err != nil {
			panic(err)
		}

		return nil
	}, func(ctx context.Context, db *bun.DB) error {
		_, err := db.NewDropTable().Model((*dao.Report)(nil)).IfExists().Exec(ctx)
		if err != nil {
			panic(err)
		}

		return nil
	})
}
