package service

import (
	"playground/playground/errors/dao"

	"github.com/pkg/errors"
)

func GetBeeInfo() {
	err := dao.GetBee()
	errors.Cause(err)
}
