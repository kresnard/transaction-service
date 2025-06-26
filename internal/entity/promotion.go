package entity

import "time"

type Promotion struct {
	ID                int    	`gorm:"primaryKey;autoIncrement:true;column:id" json:"id"`
	Name              string 	`gorm:"column:name" json:"name"`
	Type              string 	`gorm:"column:type" json:"type"`
	TargetSKU         string 	`gorm:"column:target_sku" json:"target_sku"`
	ConditionQuantity int 		`gorm:"column:condition_quantity" json:"condition_quantity"`
	DiscountPercent   float64 	`gorm:"column:discount_percent" json:"discount_percent"`
	FreeSKU           string 	`gorm:"column:free_sku" json:"free_sku"`
	CreatedAt         time.Time `gorm:"column:created_at" json:"-s"`
}