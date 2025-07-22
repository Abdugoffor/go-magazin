package order_model

import (
	auth_model "category-crud/module/auth/model"
	product_model "category-crud/module/shop/product/model"
	"time"
)

type Order struct {
	ID        int `gorm:"primarykey"`
	UserID    int `json:"user_id"`
	User      auth_model.User
	ProductID int `json:"product_id"`
	Product   product_model.Product
	Quantity  int `json:"quantity"`
	Price     int `json:"price"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Order) TableName() string {
	return "orders"
}
