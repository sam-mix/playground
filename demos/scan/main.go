package main

import (
	"fmt"
	"playground/util"

	"gorm.io/gorm"
)

type P struct {
	gorm.Model

	S uint64
	C uint64
}

type P1 struct {
	S uint64
	C uint64
}

func main() {
	db := util.Conn()

	db.Create(&P{S: 1, C: 2})

	list := make([]*P1, 0)
	db.Model(&P{}).Scan(&list)

	for _, a := range list {
		fmt.Println(a)

	}
}
