package db

import "github.com/jinzhu/gorm"

const (
	dbConfigPattern      = "host=%s port=%d user=%s dbname=%s sslmode=%s password=%s"
	PostgresqlDriverName = "postgres"
)

type Client struct {
	DBMaster *gorm.DB
}
