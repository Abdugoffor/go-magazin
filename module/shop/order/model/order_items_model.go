package order_model

import (
	product_model "category-crud/module/shop/product/model"
	"time"
)

type OrderItems struct {
	ID        int `json:"id" gorm:"primaryKey"`
	OrderID   int `json:"order_id"`
	Order     Order
	ProductID int `json:"product_id"`
	Product   product_model.Product
	Quantity  int `json:"quantity"`
	Price     int `json:"price"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (OrderItems) TableName() string {
	return "order_items"
}
