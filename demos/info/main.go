package main

import (
	"playground/model"
	"playground/util"
)

func main() {
	db := util.Conn()

	db.Save(&model.Bee{ID: 1})
}
