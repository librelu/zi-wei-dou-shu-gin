package migrations

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/zi-wei-dou-shu-gin/utils/utilerrors"
)

func NewDBMigration(host, user, password, database, sslmode, migrationFilePath, driverName string, port int32) (*migrate.Migrate, error) {
	if err := validInput(host, user, password, database, sslmode, migrationFilePath, driverName, port); err != nil {
		return nil, utilerrors.Wrap(err, "can't pass the validation")
	}
	url := getMigrationURL(host, user, password, database, sslmode, driverName, port)
	filePath := getMigrationFilePath(migrationFilePath)
	db, err := sql.Open(driverName, url)
	if err != nil {
		return nil, utilerrors.Wrap(err, "failed when connect to postgresql")
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{
		DatabaseName:     database,
		StatementTimeout: defaultTimeout,
	})
	if err != nil {
		return nil, utilerrors.Wrap(err, "failed when init driver")
	}
	m, err := migrate.NewWithDatabaseInstance(
		filePath,
		driverName,
		driver,
	)
	if err != nil {
		return nil, utilerrors.Wrap(err, "failed when new a migration client")
	}
	return m, nil
}

func validInput(host, user, password, database, sslmode, migrationFilePath, driverName string, port int32) error {
	if host == "" {
		return utilerrors.New("host can't be blank")
	}
	if user == "" {
		return utilerrors.New("user can't be blank")
	}
	if database == "" {
		return utilerrors.New("database can't be blank")
	}
	if sslmode == "" {
		return utilerrors.New("sslmode can't be blank")
	}
	if migrationFilePath == "" {
		return utilerrors.New("migrationFilePath can't be blank")
	}
	if driverName == "" {
		return utilerrors.New("driverName can't be blank")
	}
	if port <= 0 {
		return utilerrors.New("port can't be zero or negative")
	}
	return nil
}

func getMigrationURL(host, user, password, database, sslmode, driverName string, port int32) string {
	return fmt.Sprintf(
		migrationsURLPattern,
		driverName,
		user,
		password,
		host,
		port,
		database,
		sslmode,
	)
}

func getMigrationFilePath(migrationFilePath string) string {
	return fmt.Sprintf(
		migrationsFilePathPattern,
		migrationFilePath,
	)
}
