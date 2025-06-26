package repository

import (
	"transaction-service/config"
	"transaction-service/pkg/logger"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
	l *logger.Logger
	cfg *config.Config
}

func NewRepository(db *gorm.DB,l *logger.Logger, cfg *config.Config) Repository {
	return Repository{
		db: db,
		l: l,
		cfg: cfg,
	}
}