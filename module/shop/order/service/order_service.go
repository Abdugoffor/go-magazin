package order_service

import (
	"category-crud/helper"
	order_dto "category-crud/module/shop/order/dto"
	order_model "category-crud/module/shop/order/model"
	product_model "category-crud/module/shop/product/model"
	"context"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type OrderService interface {
	All(ctx echo.Context) (helper.PaginatedResponse[order_dto.OrderResponse], error)
	Show(ctx echo.Context, id uint) (order_dto.OrderResponse, error)
	Create(ctx context.Context, data order_dto.OrderCreate) (*order_dto.OrderResponse, error)
	Update(ctx context.Context, id uint, data order_dto.OrderUpdate) (*order_dto.OrderResponse, error)
	Delete(ctx echo.Context) error
}

type orderService struct {
	db *gorm.DB
}

func NewOrderService(db *gorm.DB) OrderService {
	return &orderService{
		db: db,
	}
}

func (s *orderService) All(ctx echo.Context) (helper.PaginatedResponse[order_dto.OrderResponse], error) {

	userID, err := helper.GetUserIDFromToken(ctx)

	if err != nil {
		return helper.PaginatedResponse[order_dto.OrderResponse]{}, err
	}

	var orders []order_model.Order

	result, err := helper.Paginate(ctx, s.db.Preload("OrderItems").Where("user_id = ?", userID), &orders, 10)

	if err != nil {
		return helper.PaginatedResponse[order_dto.OrderResponse]{}, err
	}

	var data []order_dto.OrderResponse
	for i := range orders {
		data = append(data, *order_dto.ToOrderResponse(&orders[i]))
	}

	return helper.PaginatedResponse[order_dto.OrderResponse]{
		Data: data,
		Meta: result.Meta,
	}, nil
}

func (s *orderService) Show(c echo.Context, id uint) (order_dto.OrderResponse, error) {
	userID, err := helper.GetUserIDFromToken(c)
	if err != nil {
		return order_dto.OrderResponse{}, err
	}

	var order order_model.Order

	if err := s.db.Preload("OrderItems").
		Where("user_id = ?", userID).
		Where("id = ?", id).
		First(&order).Error; err != nil {
		return order_dto.OrderResponse{}, err
	}

	return *order_dto.ToOrderResponse(&order), nil
}

func (s *orderService) Create(c echo.Context, data order_dto.OrderCreate) (*order_dto.OrderResponse, error) {
	userID, err := helper.GetUserIDFromToken(c)
	if err != nil {
		return nil, err
	}

	var product product_model.Product
	if err := s.db.First(&product, data.ProductID).Error; err != nil {
		return nil, err
	}

	order := order_model.Order{
		UserID:   userID,
		OrderItems: []order_model.OrderItems{
			{
				ProductID: data.ProductID,
				Quantity:  data.Quantity,
				Price:     product.Price,
			},
		},
	}

	if err := s.db.Create(&order).Error; err != nil {
		return nil, err
	}

	return order_dto.ToOrderResponse(&order), nil
}

func (s *orderService) Update(ctx context.Context, id uint, data order_dto.OrderUpdate) (*order_dto.OrderResponse, error) {
	return nil, nil
}

func (s *orderService) Delete(ctx echo.Context) error {
	return nil
}
