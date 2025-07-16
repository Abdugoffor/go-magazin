package category_service

import (
	"category-crud/helper"
	category_dto "category-crud/module/category/dto"
	category_model "category-crud/module/category/model"
	"context"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type CategoryService interface {
	All(c echo.Context) (helper.PaginatedResponse[category_dto.Response], error)
	Show(ctx context.Context, id uint) (*category_dto.Response, error)
	Create(ctx context.Context, data category_dto.Create) (*category_dto.Response, error)
	Update(ctx context.Context, id uint, data category_dto.Update) (*category_dto.Response, error)
	Delete(ctx context.Context, id uint) error
}

type categoryService struct {
	db *gorm.DB
}

func NewCategoryService(db *gorm.DB) CategoryService {
	return &categoryService{
		db: db,
	}
}

func (s *categoryService) All(ctx echo.Context) (helper.PaginatedResponse[category_dto.Response], error) {
	var categories []category_model.Category

	result, err := helper.Paginate(ctx, s.db, &categories, 10)
	if err != nil {
		return helper.PaginatedResponse[category_dto.Response]{}, err
	}

	var data []category_dto.Response
	for _, cat := range result.Data {
		data = append(data, category_dto.ToResponse(cat))
	}

	return helper.PaginatedResponse[category_dto.Response]{
		Data: data,
		Meta: result.Meta,
	}, nil
}

func (s *categoryService) Show(ctx context.Context, id uint) (*category_dto.Response, error) {
	var cat category_model.Category
	if err := s.db.WithContext(ctx).First(&cat, id).Error; err != nil {
		return nil, err
	}

	resp := category_dto.ToResponse(cat)
	return &resp, nil
}

func (s *categoryService) Create(ctx context.Context, data category_dto.Create) (*category_dto.Response, error) {
	cat := category_model.Category{
		Name:        data.Name,
		Description: data.Description,
	}
	if err := s.db.WithContext(ctx).Create(&cat).Error; err != nil {
		return nil, err
	}

	resp := category_dto.ToResponse(cat)
	return &resp, nil
}

func (s *categoryService) Update(ctx context.Context, id uint, data category_dto.Update) (*category_dto.Response, error) {
	var cat category_model.Category
	if err := s.db.WithContext(ctx).First(&cat, id).Error; err != nil {
		return nil, err
	}

	cat.Name = data.Name
	cat.Description = data.Description

	if err := s.db.WithContext(ctx).Save(&cat).Error; err != nil {
		return nil, err
	}

	resp := category_dto.ToResponse(cat)
	return &resp, nil
}

func (s *categoryService) Delete(ctx context.Context, id uint) error {
	return s.db.WithContext(ctx).Delete(&category_model.Category{}, id).Error
}
