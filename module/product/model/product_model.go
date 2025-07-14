package product_model

import (
	category_model "category-crud/module/category/model"
	"time"
)

type Product struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	CategoryID  uint   `json:"category_id"`
	Category    category_model.Category
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (Product) TableName() string {
	return "products"
}
