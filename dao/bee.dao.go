package dao

import (
	"playground/model"

	"gorm.io/gorm"
)

func List(db *gorm.DB) (bees []*model.Bee, err error) {
	err = db.Table("t_bee_with_dao").Find(&bees).Error
	return
}
