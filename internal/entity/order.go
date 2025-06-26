package entity

import "time"

type Order struct {
	ID         	int       	`gorm:"primaryKey;column:id" json:"id"`
	TotalPrice 	float64   	`gorm:"column:total_price" json:"total_price"`
	CreatedAt  	time.Time 	`gorm:"column:created_at" json:"-"`
	Items 		[]OrderItem `gorm:"foreignKey:OrderID" json:"items"`
}

type OrderItem struct {
	ID 			int 	`gorm:"primaryKey;column:id" json:"id"`
	OrderID 	int		`gorm:"column:order_id" json:"order_id"`
	SKU 		string 	`gorm:"column:sku" json:"sku"`
	Quantity	int 	`gorm:"column:quantity" json:"quantity"`
	UnitPrice	float64 `gorm:"column:unit_price" json:"unit_price"`
	SubTotal 	float64 `gorm:"column:subtotal" json:"subtotal"`
}