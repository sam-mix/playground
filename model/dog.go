package model

type Dog struct {
	ID      uint64 `gorm:"primaryKey;autoIncrement:true"`
	Name    string
	Version uint64
}
