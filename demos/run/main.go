package main

import (
	"playground/model"
	"playground/util"
)

func main() {
	db := util.Conn()

	db.First(&model.Bee{})
}
