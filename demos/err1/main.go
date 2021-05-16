package main

import (
	"playground/model"
	"playground/util"
)

func main() {
	db := util.Conn()

	db.Table("t_bee_true").Where("true = true").First(&model.Bee{})

}
