package order_dto

import (
	"category-crud/helper"
	order_model "category-crud/module/shop/order/model"
)

type OrderCreate struct {
	ProductID int `json:"product_id" validate:"required"`
	Quantity  int `json:"quantity" validate:"required"`
}

type OrderUpdate struct {
	ProductID int `json:"product_id" validate:"required"`
	Quantity  int `json:"quantity" validate:"required"`
}

type OrderResponse struct {
	ID         int                  `json:"id"`
	UserID     int                  `json:"user_id"`
	Summ       int                  `json:"summ"`
	Status     string               `json:"status"`
	OrderItems []OrdetItemsResponse `json:"order_items"`
	CreatedAt  string               `json:"created_at"`
	UpdatedAt  string               `json:"updated_at"`
}

func ToOrderResponse(order *order_model.Order) *OrderResponse {

	orderItems := make([]OrdetItemsResponse, 0)
	for _, item := range order.OrderItems {
		orderItems = append(orderItems, *ToOrderItemsResponse(&item))
	}

	return &OrderResponse{
		ID:         order.ID,
		UserID:     order.UserID,
		Summ:       order.Summ,
		Status:     order.Status,
		OrderItems: orderItems,
		CreatedAt:  helper.FormatDate(order.CreatedAt),
		UpdatedAt:  helper.FormatDate(order.UpdatedAt),
	}
}


