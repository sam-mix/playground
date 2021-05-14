package main

import (
	"playground/dao"
	"playground/util"
)

func main() {
	db := util.Conn()
	dao.List(db)
}
