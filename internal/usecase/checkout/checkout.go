package checkout

import (
	"context"
	"fmt"
	"net/http"
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
	)

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
		freeItemCount := map[string]int{}

		for sku, qty := range itemCount {
			product, err := c.repository.GetProductBySKU(ctx, tx, sku)
			if err != nil {
				return err
			}

			subTotal, freeItems, _ := c.promotion(sku, qty, product.Price, promos)

			order.Items = append(order.Items, entity.OrderItem{
				SKU: sku,
				Quantity: qty,
				UnitPrice: product.Price,
				SubTotal: float64(qty) * product.Price,
			})
			order.TotalPrice += subTotal

			for _, item := range freeItems {
				freeItemCount[item.SKU] += item.Quantity
			}
		}

		for sku, qty := range freeItemCount {
			itemCount[sku] += qty
		}

		for sku, totalQty := range itemCount {
			product, err := c.repository.GetProductBySKU(ctx, tx, sku)
			if err != nil {
				return err
			}

			if product.InventoryQty < totalQty {
				c.l.CreateLog(&logger.Log{
					Event:			"USECASE"+"|Checkout|Checkout",
					StatusCode:		http.StatusInternalServerError,
					Message: 		"product InventoryQty < totalQty",
				}, logger.LVL_ERROR)
				return fmt.Errorf("out of stock: %s", sku)
			}

			err = c.repository.ReduceInventory(ctx, tx, sku, totalQty)
			if err != nil {
				return err
			}
		}

		for sku, qty := range freeItemCount {
			order.Items = append(order.Items, entity.OrderItem{
				SKU: sku,
				Quantity: qty,
				UnitPrice: 0,
				SubTotal: 0,
			})
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

func (c Usecase) promotion(sku string, qty int, price float64, listPromo []entity.Promotion) (subTotal float64, freeItems []entity.OrderItem, promoName string ) {
	subTotal = float64(qty) * price
	freeItems = nil
	promoName = ""

	for _, promo := range listPromo {
		if  promo.TargetSKU != sku {
			continue
		}

		switch promo.Type {
		case commons.TypeBundle :
			paid := qty - (qty / promo.ConditionQuantity)
			subTotal = float64(paid) * price
			return subTotal, nil, promo.Name
		
		case commons.TypeDiscount: 
			if qty >= promo.ConditionQuantity {
				discount := price * (promo.DiscountPercent / 100)
				return float64(qty) * (price - discount), nil, promo.Name
			}

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
	
	return subTotal, nil, ""
}