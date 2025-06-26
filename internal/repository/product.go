package repository

import (
	"context"
	"net/http"
	"transaction-service/internal/entity"
	"transaction-service/pkg/logger"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (r Repository) GetProducts(ctx context.Context) (products []entity.Product, err error) {
	err = r.db.WithContext(ctx).Find(&products).Error
	if err != nil {
		r.l.CreateLog(&logger.Log{
			Event:			"REPOSITORY"+"|Product|GetProducts",
			StatusCode:		http.StatusInternalServerError,
			Message: 		"error query",
		}, logger.LVL_ERROR)
		return products, err
	}

	return
}

func (r Repository) GetProductBySKU(ctx context.Context, tx *gorm.DB, sku string) (product entity.Product, err error) {
	err = tx.WithContext(ctx).Clauses(clause.Locking{Strength: "UPDATE"}).First(&product, "sku = ?", sku).Error
	if err != nil {
		r.l.CreateLog(&logger.Log{
			Event:			"REPOSITORY"+"|Product|GetProductBySKU",
			StatusCode:		http.StatusInternalServerError,
			Message: 		"error query",
		}, logger.LVL_ERROR)
		return product, err
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
			Message: 		"error query",
		}, logger.LVL_ERROR)
		return err
	}

	return
}

func (r Repository) DB() *gorm.DB {
	return r.db
}