package repository

import (
	"context"
	"net/http"
	"transaction-service/internal/entity"
	"transaction-service/pkg/logger"
)

func (r Repository) GetPromotions(ctx context.Context) (listPromo []entity.Promotion, err error) {
	err = r.db.WithContext(ctx).Find(&listPromo).Error
	if err != nil {
		r.l.CreateLog(&logger.Log{
			Event:			"REPOSITORY"+"|Promotion|GetPromotions",
			StatusCode:		http.StatusInternalServerError,
			Message: 		"error query",
		}, logger.LVL_ERROR)
		return listPromo, err
	}
	return
}