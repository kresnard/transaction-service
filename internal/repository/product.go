package repository

import (
	"context"
	"transaction-service/internal/entity"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (r Repository) GetProducts(ctx context.Context) (products []entity.Product, err error) {
	err = r.db.WithContext(ctx).Find(&products).Error
	return
}

func (r Repository) GetProductBySKU(ctx context.Context, tx *gorm.DB, sku string) (product entity.Product, err error) {
	err = tx.WithContext(ctx).Clauses(clause.Locking{Strength: "UPDATE"}).First(&product, "sku = ?", sku).Error
	return
}

func (r Repository) ReduceInventory(ctx context.Context, tx *gorm.DB, sku string, qty int) (err error) {
	err = tx.WithContext(ctx).
	Model(&entity.Product{}).
	Where("sku = ? AND inventory_qty >= ?", sku, qty).
	Update("inventory_qty", gorm.Expr("inventory_qty - ?", qty)).Error
	return
}

func (r Repository) DB() *gorm.DB {
	return r.db
}