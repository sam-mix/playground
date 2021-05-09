package util

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"playground/myzap"
)

func Conn() *gorm.DB {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:1qaz@wsx@tcp(127.0.0.1:53306)/x?charset=utf8mb4&parseTime=True&loc=Local"

	zapLogger := &myzap.Logger{
		ZapLogger:                 myzap.NewLogger().Logger,
		LogLevel:                  logger.Info,
		SlowThreshold:             100 * time.Millisecond,
		SkipCallerLookup:          false,
		IgnoreRecordNotFoundError: true,
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         255,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		Logger: zapLogger,
	})

	if err != nil {
		panic(err)
	}

	db.Callback().Query().Before("gorm:query").Register("gorm:auto_migrate", migrate)
	db.Callback().Update().Before("gorm:update").Register("gorm:auto_migrate", migrate)
	db.Callback().Create().Before("gorm:create").Register("gorm:auto_migrate", migrate)

	return db

}

func migrate(db *gorm.DB) {
	db.AutoMigrate(db.Statement.Model)
	db.Model(nil)
}
