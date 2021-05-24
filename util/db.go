package util

import (
	"context"

	"gorm.io/gorm"
)

// CloneDB 只会复制 db.Statement
func CloneDB(db *gorm.DB) *gorm.DB {
	// return db.WithContext(db.Statement.Context)
	return db.WithContext(context.Background())
}

// NewDB 完全新创建一个 DB
func NewDB(db *gorm.DB) *gorm.DB {
	return db.Session(&gorm.Session{NewDB: true, Context: context.Background()})
	// return db.Session(&gorm.Session{NewDB: true, Context: db.Statement.Context})
}

// CleanDB 使用原有 context 创建一个新的 DB
func CleanDB(db *gorm.DB) *gorm.DB {
	// return db.Session(&gorm.Session{NewDB: true, Context: context.Background()})
	return db.Session(&gorm.Session{NewDB: true, Context: db.Statement.Context})
}
