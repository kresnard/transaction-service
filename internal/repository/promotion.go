package repository

import (
	"context"
	"errors"
	"net/http"
	"transaction-service/commons"
	"transaction-service/internal/entity"
	"transaction-service/pkg/logger"

	"gorm.io/gorm"
)

func (r Repository) GetPromotions(ctx context.Context) (listPromo []entity.Promotion, err error) {
	err = r.db.WithContext(ctx).Find(&listPromo).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			r.l.CreateLog(&logger.Log{
				Event:			"REPOSITORY"+"|Promotion|GetPromotions",
				StatusCode:		http.StatusNotFound,
				Message: 		err.Error(),
			}, logger.LVL_ERROR)
			return listPromo, commons.ErrNotFound
		}
		r.l.CreateLog(&logger.Log{
			Event:			"REPOSITORY"+"|Promotion|GetPromotions",
			StatusCode:		http.StatusInternalServerError,
			Message: 		err.Error(),
		}, logger.LVL_ERROR)
		return listPromo, commons.ErrFailedGetData
	}

	return
}