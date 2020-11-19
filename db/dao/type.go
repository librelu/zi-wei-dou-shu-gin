package dao

import (
	"github.com/zi-wei-dou-shu-gin/db"
)

type Dao struct {
	dbClient *db.Client
}

type DaoHandler interface {
	SaveBoard(gender, name string, birthday int64) error
}

type Board struct {
	Name     string
	Gender   string
	Birthday int64
}
