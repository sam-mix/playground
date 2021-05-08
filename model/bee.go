package model

type Bee struct {
	ID      uint64 `gorm:"primaryKey"`
	Name    string
	Version uint64
}
