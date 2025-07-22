package order_dto

import (
	auth_dto "category-crud/module/auth/dto"
	order_model "category-crud/module/shop/order/model"
	product_dto "category-crud/module/shop/product/dto"
)

type OrderCreate struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type OrderUpdate struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type OrderResponse struct {
	ID       int                   `json:"id"`
	User     auth_dto.UserResponse `json:"user"`
	Product  product_dto.Response  `json:"product"`
	Quantity int                   `json:"quantity"`
	Price    int                   `json:"price"`
}

func ToOrderResponse(order order_model.Order) OrderResponse {
	return OrderResponse{
		ID:       order.ID,
		User:     auth_dto.ToUserResponse(order.User),
		Product:  product_dto.ToResponse(order.Product),
		Quantity: order.Quantity,
		Price:    order.Price,
	}
}
