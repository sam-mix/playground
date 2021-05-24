package main

import (
	"errors"
	"fmt"
	"playground/model"
	"playground/util"

	"gorm.io/gorm"
)

func main() {
	db := util.Sqlite()
	d := &model.Dog{Name: "bee1"}
	if err := db.Table("t_dog_002").Create(d).Error; err != nil {
		panic(err)
	}

	if err := db.Table("t_dog_002").Where("true = true AND id = ?", d.ID).First(&model.Dog{}).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}

	d1 := &model.Bee{}
	if err := db.Table("t_dog_002").Where("1 = 1 AND id = ?", d.ID).First(d1).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}

	d2 := &model.Bee{}
	if err := db.Table("t_dog_002").Where("`1` = `1` AND id = ?", d.ID).First(d2).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}
	fmt.Println(d2)
}
