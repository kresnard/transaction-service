package repository

import (
	"context"
	"errors"
	"net/http"
	"transaction-service/commons"
	"transaction-service/internal/entity"
	"transaction-service/pkg/logger"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (r Repository) GetProducts(ctx context.Context) (products []entity.Product, err error) {
	err = r.db.WithContext(ctx).Find(&products).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			r.l.CreateLog(&logger.Log{
				Event:			"REPOSITORY"+"|Product|GetProducts",
				StatusCode:		http.StatusNotFound,
				Message: 		err.Error(),
			}, logger.LVL_ERROR)
			return products, commons.ErrNotFound
		}
		r.l.CreateLog(&logger.Log{
			Event:			"REPOSITORY"+"|Product|GetProducts",
			StatusCode:		http.StatusInternalServerError,
			Message: 		err.Error(),
		}, logger.LVL_ERROR)
		return products, commons.ErrFailedGetData
	}

	return
}

func (r Repository) GetProductBySKU(ctx context.Context, tx *gorm.DB, sku string) (product entity.Product, err error) {
	err = tx.WithContext(ctx).Clauses(clause.Locking{Strength: "UPDATE"}).First(&product, "sku = ?", sku).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			r.l.CreateLog(&logger.Log{
				Event:			"REPOSITORY"+"|Product|GetProductBySKU",
				StatusCode:		http.StatusNotFound,
				Message: 		err.Error(),
			}, logger.LVL_ERROR)
			return product, commons.ErrNotFound
		}
		r.l.CreateLog(&logger.Log{
			Event:			"REPOSITORY"+"|Product|GetProductBySKU",
			StatusCode:		http.StatusInternalServerError,
			Message: 		err.Error(),
		}, logger.LVL_ERROR)
		return product, commons.ErrFailedGetData
	}


	return
}

func (r Repository) ReduceInventory(ctx context.Context, tx *gorm.DB, sku string, qty int) (err error) {
	err = tx.WithContext(ctx).
	Model(&entity.Product{}).
	Where("sku = ? AND inventory_qty >= ?", sku, qty).
	Update("inventory_qty", gorm.Expr("inventory_qty - ?", qty)).Error

	if err != nil {
		r.l.CreateLog(&logger.Log{
			Event:			"REPOSITORY"+"|Product|ReduceInventory",
			StatusCode:		http.StatusInternalServerError,
			Message: 		err.Error(),
		}, logger.LVL_ERROR)
		return commons.ErrUpdateData
	}

	return
}

func (r Repository) DB() *gorm.DB {
	return r.db
}