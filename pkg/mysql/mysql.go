package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"time"
	"transaction-service/config"
	"transaction-service/pkg/logger"

	_ "github.com/go-sql-driver/mysql"
)

func New(driverName string, cfg *config.Config, l *logger.Logger) *sql.DB {
	db, err := sql.Open("mysql", cfg.MYSQL.URL)
	if err != nil {
		log.Fatalf(fmt.Sprintf("not success to connect to database: %v", err))
	}

	db.SetMaxIdleConns(cfg.MYSQL.MaxIdleConns)
	db.SetMaxOpenConns(cfg.MYSQL.MaxOpenConns)
	db. SetConnMaxLifetime(time.Duration(cfg.MYSQL.MaxLifeTimeConns) *  time.Second)
	return db
}