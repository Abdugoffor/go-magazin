package product_dto

import (
	"category-crud/helper"
	category_dto "category-crud/module/category/dto"
	product_model "category-crud/module/product/model"
)

type Create struct {
	Name        string `json:"name" validate:"required,max=255"`
	Description string `json:"description" validate:"required,max=255"`
	Price       int    `json:"price" validate:"required"`
	CategoryID  uint   `json:"category_id" validate:"required"`
}

type Update struct {
	Name        string `json:"name" validate:"required,max=255"`
	Description string `json:"description" validate:"required,max=255"`
	Price       int    `json:"price" validate:"required"`
	CategoryID  uint   `json:"category_id" validate:"required"`
}

type Response struct {
	ID          uint                              `json:"id"`
	Name        string                            `json:"name"`
	Description string                            `json:"description"`
	Price       int                               `json:"price"`
	Category    category_dto.ResponseWithProducts `json:"category"`
	CreatedAt   string                            `json:"created_at"`
	UpdatedAt   string                            `json:"updated_at"`
}

func ToResponse(product product_model.Product) Response {
	return Response{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Category: category_dto.ResponseWithProducts{
			ID:   product.Category.ID,
			Name: product.Category.Name,
			// Description: product.Category.Description,
		},
		CreatedAt: helper.FormatDate(product.CreatedAt),
		UpdatedAt: helper.FormatDate(product.UpdatedAt),
	}
}
