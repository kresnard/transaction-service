package checkout

import (
	"context"
	"errors"
	"fmt"
	"transaction-service/commons"
	"transaction-service/internal/entity"

	"transaction-service/config"

	"transaction-service/internal/repository"
	"transaction-service/pkg/logger"

	"gorm.io/gorm"
)

type Usecase struct {
	repository repository.IRepository
	l *logger.Logger
	cfg *config.Config

}

func NewUsecase(repository repository.IRepository,	l *logger.Logger,	cfg *config.Config) Usecase {
	return Usecase{
		repository: repository,
		l: l,
		cfg: cfg,
	}
}

type ICheckoutUsecase interface {
	Checkout(ctx context.Context, skus []string) (*entity.Order, error)
}

func (c Usecase) Checkout(ctx context.Context, skus []string) (res *entity.Order, err error) {
	var ( 
		order entity.Order
		orderItem entity.OrderItem
	)
	
	if len(skus) == 0 {
		return res, errors.New("empty cart")
	}

	itemCount := map[string]int{}
	for _, sku := range skus {
		itemCount[sku]++
	}

	promos, err := c.repository.GetPromotions(ctx)
	if err != nil {
		return res, err
	}

	db :=  c.repository.DB()
	

	err = db.Transaction(func(tx *gorm.DB) error {
		for sku, qty := range itemCount {
			product, err := c.repository.GetProductBySKU(ctx, tx, sku)
			if err != nil {
				return err
			}

			if product.InventoryQty < qty {
				return fmt.Errorf("out of stock: %s", sku)
			}

			err = c.repository.ReduceInventory(ctx, tx, sku, qty)
			if err != nil {
				return err
			}

			subTotal, freeItems, _ := c.promotion(sku, qty, product.Price, promos)
			orderItem.SKU = sku
			orderItem.UnitPrice = product.Price
			orderItem.Quantity = qty
			orderItem.SubTotal = subTotal
			order.Items = append(order.Items, orderItem)
			order.Items = append(order.Items, freeItems...)
			order.TotalPrice += subTotal
		}

		err := c.repository.Save(ctx, tx, &order)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return res, err
	}

	return &order, nil
}

func (c Usecase) promotion(sku string, qty int, price float64, listPromo []entity.Promotion) (subTotal float64, orderItems []entity.OrderItem, promoName string ) {
	for _, promo := range listPromo {
		if  promo.TargetSKU != sku || qty < promo.ConditionQuantity {
			continue
		}

		switch promo.Type {
		case commons.TypeBundle :
			paid := qty - (qty / promo.ConditionQuantity)
			return float64(paid) * price, nil, promo.Name
		
		case commons.TypeDiscount: 
			discount := price * (promo.DiscountPercent / 100)
			return float64(qty) * (price - discount), nil, promo.Name

		case commons.TypeFreebie:
			freeItems := []entity.OrderItem{
				{
					SKU: promo.FreeSKU,
					Quantity: qty,
					UnitPrice: 0,
					 SubTotal: 0,
				},
			}
			return float64(qty)*price, freeItems, promo.Name
		}
	}
	
	return
}