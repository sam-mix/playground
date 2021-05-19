package main

import (
	"fmt"
	"playground/util"
)

type NewInsert struct {
	ID       uint `gorm:"primaryKey;autoIncrement:true"`
	Name     string
	UpdateAt uint64
}

func main() {
	db := util.Conn()
	newInsert := &NewInsert{Name: "Paul"}
	if err := db.Table("t_new_insert").Omit("NewInsert.UpdateAt").Create(newInsert).Error; err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", newInsert)
}
