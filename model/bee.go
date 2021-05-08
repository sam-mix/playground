package model

type Bee struct {
	ID      uint64 `gorm:"primaryKey;autoIncrement:false"`
	Name    string
	Version uint64
}
