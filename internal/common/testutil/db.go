package testutil

import (
	"log"
	"testing"

	"github.com/citcho/go-gizlog/internal/common/config"
	"github.com/citcho/go-gizlog/internal/common/database"
	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
)

func OpenDBForTest(t *testing.T) *bun.DB {
	t.Helper()

	cfg, err := config.NewDBConfig()
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	db := database.NewDB(cfg)

	t.Cleanup(
		func() { _ = db.Close() },
	)

	return db
}
