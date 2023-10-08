package database

import (
	"database/sql"
	"fmt"

	"github.com/citcho/go-gizlog/internal/common/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

func NewDB(cfg config.DBConfig) *bun.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	sqldb, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("aaa")
		panic(err)
	}

	return bun.NewDB(sqldb, mysqldialect.New())
}
