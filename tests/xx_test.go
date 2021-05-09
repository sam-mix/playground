package tests

import (
	"playground/model"
	"playground/util"
	"testing"
)

func TestXXX(t *testing.T) {
	db := util.Conn()
	t.Run("xxx", func(t *testing.T) {
		db.Save(&model.Bee{ID: 1})
	})
}
