package migrator

import (
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

func Up(pool *pgxpool.Pool, migrationsDir string) error {
	db := stdlib.OpenDBFromPool(pool)

	err := goose.Up(db, migrationsDir)
	if err != nil {
		if !errors.Is(err, goose.ErrNoMigrations) && !errors.Is(err, goose.ErrNoNextVersion) {
			return nil
		}
	}

	return nil
}
