package util

import (
	"fmt"
	"time"

	"code.guanmai.cn/back_end/grpckit/config"
	"google.golang.org/grpc/grpclog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Open(conf *config.DatabaseConfig) (*gorm.DB, error) {
	database, ok := databases[conf.Name]
	if !ok {
		mutex.Lock()
		defer mutex.Unlock()

		database, ok = databases[conf.Name]
		if !ok {
			database = &Database{}

			var parameters string
			for name, value := range conf.Parameters {
				if parameters == "" {
					parameters = fmt.Sprintf("%s=%s", name, value)
				} else {
					parameters += fmt.Sprintf("&%s=%s", name, value)
				}
			}
			dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s&charset=utf8mb4", conf.User, conf.Password, conf.Host, conf.Port, conf.Name, parameters)

			configLogger := logging.ConfigLogger()
			zapLogger := logging.Logger{
				ZapLogger:                 configLogger.Logger,
				LogLevel:                  logger.Warn,
				SlowThreshold:             100 * time.Millisecond,
				SkipCallerLookup:          false,
				IgnoreRecordNotFoundError: true,
			}
			zapLogger.SetAsDefault() // optional: configure gorm to use this zapgorm.Logger for callbacks

			// zapLogger := logger.New(
			// 	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			// 	logger.Config{
			// 		SlowThreshold:             time.Second,   // Slow SQL threshold
			// 		LogLevel:                  logger.Silent, // Log level
			// 		IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			// 		Colorful:                  true,          // Disable color
			// 	},
			// )
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
				grpclog.Errorf("open database %s error: %v", err, conf.Name)
				return nil, err
			}

			if conf.Driver == "mysql" {
				db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci")
			}

			db.Callback().Create().Before("gorm:create").Register("gorm:auto_migrate", database.migrate)
			db.Callback().Update().Before("gorm:update").Register("gorm:auto_migrate", database.migrate)
			db.Callback().Query().Before("gorm:query").Register("gorm:auto_migrate", database.migrate)
			// db.Callback().Delete().Before("gorm:delete").Register("gorm:auto_migrate", database.migrate)
			// db.Callback().Raw().Before("gorm:raw").Register("gorm:auto_migrate", database.migrate)
			// db.Callback().Row().Before("gorm:row").Register("gorm:auto_migrate", database.migrate)

			sqlDB, err := db.DB()
			if err != nil {
				return nil, err
			}
			if conf.MaxOpenConns != 0 {
				sqlDB.SetMaxOpenConns(int(conf.MaxOpenConns))
			}
			if conf.MaxIdleConns != 0 {
				sqlDB.SetMaxIdleConns(int(conf.MaxIdleConns))
			}

			if conf.Debug || true {
				database.db = db.Debug()
			} else {
				database.db = db
			}
			grpclog.Infof("open database %s", conf.Name)

			databases[conf.Name] = database
		}
	}

	return database.db, nil
}

func (db *Database) migrate(gormDB *gorm.DB) {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	table := gormDB.Statement.Table
	if _, ok := db.tables.Load(table); !ok {
		if _, ok := db.tables.Load(table); !ok {
			value := gormDB.Statement.Model
			grpclog.Infof("migrate table %s %v", table, value)
			gormDB.AutoMigrate(value)
			db.tables.Store(table, true)
		}
	}
	gormDB.Model(nil)
}
