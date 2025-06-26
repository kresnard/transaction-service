package repository

import (
	"context"
	"transaction-service/internal/entity"
)

func (r Repository) GetPromotions(ctx context.Context) (listPromo []entity.Promotion, err error) {
	err = r.db.WithContext(ctx).Find(&listPromo).Error
	return
}