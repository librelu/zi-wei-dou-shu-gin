package configs

import (
	"github.com/spf13/viper"
	"github.com/zi-wei-dou-shu-gin/utils/utilerrors"
)

func NewDBConfig() (*DBConfig, error) {
	dbConfig := viper.Sub("db")
	if dbConfig == nil {
		return nil, utilerrors.New("db can't be an empty object")
	}
	host := dbConfig.GetString("host")
	if host == "" {
		return nil, utilerrors.New("host can't be blank in config")
	}

	port := dbConfig.GetInt32("port")
	if port <= 0 {
		return nil, utilerrors.New("port can't be zero or negative in config")
	}

	user := dbConfig.GetString("user")
	if user == "" {
		return nil, utilerrors.New("user can't be blank in config")
	}

	database := dbConfig.GetString("database")
	if database == "" {
		return nil, utilerrors.New("database can't be blank in config")
	}

	sslMode := dbConfig.GetString("sslmode")
	if sslMode == "" {
		return nil, utilerrors.New("sslmode can't be blank in config")
	}

	return &DBConfig{
		Host:     host,
		Port:     port,
		User:     user,
		Password: dbConfig.GetString("password"),
		Database: database,
		SSLMode:  sslMode,
	}, nil
}
