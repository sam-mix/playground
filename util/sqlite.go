package util

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"playground/myzap"
)

func Sqlite() *gorm.DB {

	zapLogger := &myzap.Logger{
		ZapLogger:                 myzap.NewLogger().Logger,
		LogLevel:                  logger.Info,
		SlowThreshold:             100 * time.Millisecond,
		SkipCallerLookup:          false,
		IgnoreRecordNotFoundError: true,
	}

	// github.com/mattn/go-sqlite3
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{Logger: zapLogger})
	if err != nil {
		panic(err)
	}

	db.Callback().Query().Before("gorm:query").Register("gorm:auto_migrate", migrate)
	db.Callback().Update().Before("gorm:update").Register("gorm:auto_migrate", migrate)
	db.Callback().Create().Before("gorm:create").Register("gorm:auto_migrate", migrate)

	return db

}
