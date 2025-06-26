package repository

import (
	"context"
	"net/http"
	"transaction-service/internal/entity"
	"transaction-service/pkg/logger"

	"gorm.io/gorm"
)

func (r Repository) Save(ctx context.Context, tx *gorm.DB, order *entity.Order) (err error) {
	err = tx.WithContext(ctx).Create(order).Error
	if err != nil {
		r.l.CreateLog(&logger.Log{
			Event:			"REPOSITORY"+"|Order|Save",
			StatusCode:		http.StatusInternalServerError,
			Request: 		order,
			Message: 		"error query",
		}, logger.LVL_ERROR)
		return err
	}

	return nil
}