package db

import (
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"sf-go/logs"
	"time"
)

type DB struct {
	ReadDB  *gorm.DB
	WriteDB *gorm.DB
}

func Open(dbCfg DatabaseConfig, dst ...any) (*gorm.DB, error) {
	// 链接数据库
	db, err := gorm.Open(mysql.Open(dbCfg.GetDNS()), &gorm.Config{
		Logger: logger.New(
			logrus.StandardLogger(),
			logger.Config{
				SlowThreshold:             3 * time.Second,
				LogLevel:                  logger.Error,
				IgnoreRecordNotFoundError: true,
				Colorful:                  false,
			},
		),
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		return nil, err
	}
	// 设置数据库连接池参数
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(dbCfg.MaxOpenConns)
	sqlDB.SetMaxIdleConns(dbCfg.MaxIdleConns)
	sqlDB.SetConnMaxIdleTime(time.Second * 30)
	sqlDB.SetConnMaxLifetime(time.Minute)

	if len(dst) == 0 {
		return db, nil
	}

	if err = db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(dst...); err != nil {
		return nil, err
	}
	return db, nil
}

func NewDB(
	wcfg DatabaseConfig,
	rcfg DatabaseConfig,
	dst ...any,
) (*DB, error) {
	var (
		err error
	)
	readDB, err := Open(rcfg, dst...)
	if err != nil {
		logs.Logger.Error("Open ReaderDatabase Error", zap.Error(err))
		return nil, err
	}
	writeDB, err := Open(wcfg, dst...)
	if err != nil {
		logs.Logger.Error("Open WriterDatabase Error", zap.Error(err))
		return nil, err
	}
	logs.Logger.Info("Open Database Ok")
	return &DB{
		ReadDB:  readDB,
		WriteDB: writeDB,
	}, nil
}
