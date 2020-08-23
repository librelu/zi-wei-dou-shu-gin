package cmd

import "github.com/golang-migrate/migrate/v4"

var (
	ginMode         string
	configPath      string
	fileType        string
	migrationPath   string
	migrationDriver string
	steps           int
	dbMigration     *migrate.Migrate
)
