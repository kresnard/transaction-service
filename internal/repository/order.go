package repository

import (
	"context"
	"transaction-service/internal/entity"

	"gorm.io/gorm"
)

func (r Repository) Save(ctx context.Context, tx *gorm.DB, order *entity.Order) (err error) {
	err = tx.WithContext(ctx).Create(order).Error
	return
}