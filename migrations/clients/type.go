package migrations

const (
	migrationsURLPattern      = "%s://%s:%s@%s:%d/%s?sslmode=%s"
	migrationsFilePathPattern = "file://%s"
	defaultTimeout            = 1000
	PostgresDriverName        = "postgres"
	DefaultMigrationPath      = "./migrations/data/"
	DefaultConfigPath         = "./configs/yaml"
	YamlType                  = "yaml"
)
