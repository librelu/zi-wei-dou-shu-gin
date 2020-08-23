package db

import (
	"fmt"

	"github.com/bearners-gin/utils/utilerrors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func NewDBClient(host, user, password, database, sslmode string, port int32) (client *Client, err error) {
	if host == "" {
		return nil, utilerrors.New("host can't be nil when new db client")
	}
	if user == "" {
		return nil, utilerrors.New("user can't be nil when new db client")
	}
	if database == "" {
		return nil, utilerrors.New("database can't be nil when new db client")
	}
	if sslmode == "" {
		return nil, utilerrors.New("sslmode can't be nil when new db client")
	}
	if port < 0 {
		return nil, utilerrors.New("sslmode can't be zero or negative")
	}

	config := getConfig(host, user, password, database, sslmode, port)
	db, err := gorm.Open(PostgresqlDriverName, config)
	if err != nil {
		return nil, utilerrors.Wrap(err, "connect to db failed when new db client")
	}

	return &Client{
		DBMaster: db,
	}, nil
}

func getConfig(host, user, password, database, sslmode string, port int32) string {
	return fmt.Sprintf(
		dbConfigPattern, host, port, user, database, sslmode, password,
	)
}
