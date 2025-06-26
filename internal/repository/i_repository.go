package repository

import (
	"context"
	"transaction-service/internal/entity"

	"gorm.io/gorm"
)

type IRepository interface {
	ProductRepository
	OrderRepository
	PromotionRepository
}

type ProductRepository interface {
	GetProducts(ctx context.Context) ([]entity.Product, error)
	GetProductBySKU(ctx context.Context, tx *gorm.DB, sku string) (entity.Product, error)
	ReduceInventory(ctx context.Context, tx *gorm.DB, sku string, qty int) error
	DB() *gorm.DB
}

type OrderRepository interface {
	Save(ctx context.Context, tx *gorm.DB, order *entity.Order) error
}

type PromotionRepository interface {
	GetPromotions(ctx context.Context) ([]entity.Promotion, error)
}