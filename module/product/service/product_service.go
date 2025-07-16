package product_service

import (
	"category-crud/helper"
	product_dto "category-crud/module/product/dto"
	product_model "category-crud/module/product/model"
	"context"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ProductService interface {
	All(ctx echo.Context) (helper.PaginatedResponse[product_dto.Response], error)
	Show(ctx context.Context, id uint) (*product_dto.Response, error)
	Create(ctx context.Context, data product_dto.Create) (*product_dto.Response, error)
	Update(ctx context.Context, id uint, data product_dto.Update) (*product_dto.Response, error)
	Delete(ctx context.Context, id uint) error
}

type productService struct {
	db *gorm.DB
}

func NewProductService(db *gorm.DB) ProductService {
	return &productService{
		db: db,
	}
}

func (s *productService) All(ctx echo.Context) (helper.PaginatedResponse[product_dto.Response], error) {
	var products []product_model.Product

	result, err := helper.Paginate(ctx, s.db.Preload("Category"), &products, 10)
	if err != nil {
		return helper.PaginatedResponse[product_dto.Response]{}, err
	}

	var data []product_dto.Response
	for _, product := range result.Data {
		data = append(data, product_dto.ToResponse(product))
	}

	return helper.PaginatedResponse[product_dto.Response]{
		Data: data,
		Meta: result.Meta,
	}, nil
}

func (s *productService) Show(ctx context.Context, id uint) (*product_dto.Response, error) {
	var product product_model.Product

	if err := s.db.WithContext(ctx).Preload("Category").First(&product, id).Error; err != nil {
		return nil, err
	}

	resp := product_dto.ToResponse(product)

	return &resp, nil
}

func (s *productService) Create(ctx context.Context, data product_dto.Create) (*product_dto.Response, error) {
	product := product_model.Product{
		Name:        data.Name,
		Description: data.Description,
		Price:       data.Price,
		CategoryID:  data.CategoryID,
	}

	if err := s.db.WithContext(ctx).Preload("Category").Create(&product).Error; err != nil {
		return nil, err
	}

	if err := s.db.WithContext(ctx).Preload("Category").First(&product, product.ID).Error; err != nil {
		return nil, err
	}

	resp := product_dto.ToResponse(product)
	return &resp, nil
}

func (s *productService) Update(ctx context.Context, id uint, data product_dto.Update) (*product_dto.Response, error) {
	var product product_model.Product

	if err := s.db.WithContext(ctx).First(&product, id).Error; err != nil {
		return nil, err
	}

	product.Name = data.Name
	product.Description = data.Description
	product.Price = data.Price
	product.CategoryID = data.CategoryID

	if err := s.db.WithContext(ctx).Preload("Category").Save(&product).Error; err != nil {
		return nil, err
	}

	if err := s.db.WithContext(ctx).Preload("Category").First(&product, product.ID).Error; err != nil {
		return nil, err
	}

	resp := product_dto.ToResponse(product)
	return &resp, nil
}

func (s *productService) Delete(ctx context.Context, id uint) error {
	return s.db.WithContext(ctx).Delete(&product_model.Product{}, id).Error
}
