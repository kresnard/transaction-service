package entity

type Product struct {
	SKU          string  `gorm:"primaryKey;column:sku" json:"sku"`
	Name         string  `gorm:"column:name" json:"name"`
	Price        float64 `gorm:"column:price" json:"price"`
	InventoryQty int     `gorm:"column:inventory_qty" json:"inventory_qty"`
}