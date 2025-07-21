package order_dto

import (
	order_model "category-crud/module/shop/order/model"
	product_dto "category-crud/module/shop/product/dto"
)

type OrdetItemsCreate struct {
	OrderID   int `json:"order_id" validate:"required"`
	ProductID int `json:"product_id" validate:"required"`
	Quantity  int `json:"quantity" validate:"required"`
}

type OrdetItemsUpdate struct {
	OrderID   int `json:"order_id" validate:"required"`
	ProductID int `json:"product_id" validate:"required"`
	Quantity  int `json:"quantity" validate:"required"`
}

type OrdetItemsResponse struct {
	ID       int                           `json:"id"`
	Product  product_dto.ResponseWithOrder `json:"product"`
	Quantity int                           `json:"quantity"`
	Price    int                           `json:"price"`
}

func ToOrderItemsResponse(orderItem *order_model.OrderItems) *OrdetItemsResponse {
	return &OrdetItemsResponse{
		ID: orderItem.ID,
		Product: product_dto.ResponseWithOrder{
			Name: orderItem.Product.Name,
		},
		Quantity: orderItem.Quantity,
		Price:    orderItem.Price,
	}
}
