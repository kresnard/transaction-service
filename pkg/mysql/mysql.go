package mysql

import (
	"log"
	"time"
	"transaction-service/config"
	loggerTrx "transaction-service/pkg/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)
func New(cfg *config.Config, l *loggerTrx.Logger) *gorm.DB {
	dsn := cfg.MYSQL.URL
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get generic DB from GORM: %v", err)
	}

	sqlDB.SetMaxIdleConns(cfg.MYSQL.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MYSQL.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.MYSQL.MaxLifeTimeConns) * time.Second)

	return db
}