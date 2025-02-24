package main

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func MigratePG(pg *sql.DB) error {
	driver, err := postgres.WithInstance(pg, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance("file://./migrations", "postgres", driver)
	if err != nil {
		return err
	}

	mErr := m.Up()
	if mErr != nil{
		if mErr.Error() != "no change"{
			return mErr
		}
	}

	return nil
}