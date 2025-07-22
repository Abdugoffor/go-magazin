package order_service

import (
	"category-crud/helper"
	order_dto "category-crud/module/shop/order/dto"
	order_model "category-crud/module/shop/order/model"
	product_model "category-crud/module/shop/product/model"
	"fmt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type OrderService interface {
	All(c echo.Context) (helper.PaginatedResponse[order_dto.OrderResponse], error)
	Show(c echo.Context, id uint) (*order_dto.OrderResponse, error)
	Create(c echo.Context, data order_dto.OrderCreate) (*order_dto.OrderResponse, error)
	Update(c echo.Context, id uint, data order_dto.OrderUpdate) (*order_dto.OrderResponse, error)
	Delete(c echo.Context, id uint) error
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
	var models []order_model.Order

	query := s.db.Preload("User").Preload("Product")

	result, err := helper.Paginate(ctx, query, &models, 10)

	if err != nil {
		return helper.PaginatedResponse[order_dto.OrderResponse]{}, err
	}

	var data []order_dto.OrderResponse
	for _, model := range result.Data {
		data = append(data, order_dto.ToOrderResponse(model))
	}

	return helper.PaginatedResponse[order_dto.OrderResponse]{
		Data: data,
		Meta: result.Meta,
	}, nil
}

func (s *orderService) Show(c echo.Context, id uint) (*order_dto.OrderResponse, error) {
	var model []order_model.Order

	if err := s.db.Where("id = ?", id).Preload("User").Preload("Product").Find(&model).Error; err != nil {
		return nil, err
	}

	resp := order_dto.ToOrderResponse(model[0])

	return &resp, nil
}

func (s *orderService) Create(c echo.Context, data order_dto.OrderCreate) (*order_dto.OrderResponse, error) {

	userID, err := helper.GetUserFromToken(c)

	if err != nil {
		return nil, fmt.Errorf("user_id olishda xatolik: %v", err)
	}

	var product product_model.Product
	if err := s.db.First(&product, data.ProductID).Error; err != nil {
		return nil, fmt.Errorf("product not found")
	}

	totalPrice := product.Price * int(data.Quantity)

	order := order_model.Order{
		UserID:    int(userID),
		ProductID: data.ProductID,
		Quantity:  int(data.Quantity),
		Price:     totalPrice,
	}

	if err := s.db.Create(&order).Error; err != nil {
		return nil, fmt.Errorf("failed to create order: %v", err)
	}

	if err := s.db.Preload("User").Preload("Product").First(&order, order.ID).Error; err != nil {
		return nil, fmt.Errorf("failed to load order details: %v", err)
	}

	response := order_dto.ToOrderResponse(order)
	return &response, nil
}

func (s *orderService) Update(c echo.Context, id uint, data order_dto.OrderUpdate) (*order_dto.OrderResponse, error) {
	return nil, nil
}

func (s *orderService) Delete(c echo.Context, id uint) error {
	return nil
}
