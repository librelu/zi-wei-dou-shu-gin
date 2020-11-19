package dao

import (
	"github.com/zi-wei-dou-shu-gin/db"
)

func NewDao(dbClient *db.Client) DaoHandler {
	return &Dao{
		dbClient: dbClient,
	}
}

func (d *Dao) SaveBoard(gender, name string, birthday int64) error {
	if err := d.dbClient.DBMaster.Create(&Board{
		Name:     name,
		Gender:   gender,
		Birthday: birthday,
	}).Error; err != nil {
		return err
	}
	return nil
}
