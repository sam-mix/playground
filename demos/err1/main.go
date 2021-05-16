package main

import (
	"errors"
	"playground/model"
	"playground/util"

	"gorm.io/gorm"
)

func main() {
	db := util.Conn()

	// db.Table("t_bee_true").Where("true = true").First(&model.Bee{})
	// db.Table("t_bee_true").Where("1 = 1").First(&model.Bee{})
	if err := db.Table("t_bee_true").Where("1 = 1").First(&model.Bee{}).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		panic(err)
	}
}
