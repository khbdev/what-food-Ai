package migration

import (
	"user-product-service/internal/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations() {
m, err := migrate.New(
    "file://user-product-service/migrations",
    config.PostgresURL(),
)

	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil && err.Error() != "no change" {
		panic(err)
	}
}