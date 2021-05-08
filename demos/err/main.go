package main

import (
	"playground/model"
	"playground/util"
)

func main() {
	db := util.Conn()

	db.Save(&model.Bee{})
	db.Save(&model.Bee{})
}
