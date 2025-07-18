package auth_service

import (
	"category-crud/helper"
	auth_dto "category-crud/module/auth/dto"
	auth_model "category-crud/module/auth/model"
	"context"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type PermissionService interface {
	All(ctx echo.Context) (helper.PaginatedResponse[auth_dto.PermissionResponse], error)
	Show(ctx context.Context, id uint) (*auth_dto.PermissionResponse, error)
	Update(ctx context.Context, id uint, data auth_dto.UpdatePermission) (*auth_dto.PermissionResponse, error)
}

type permissionService struct {
	db *gorm.DB
}

func NewPermissionService(db *gorm.DB) PermissionService {
	return &permissionService{
		db: db,
	}
}

func (s *permissionService) All(ctx echo.Context) (helper.PaginatedResponse[auth_dto.PermissionResponse], error) {
	var permissions []auth_model.Permission

	result, err := helper.Paginate(ctx, s.db, &permissions, 10)
	if err != nil {
		return helper.PaginatedResponse[auth_dto.PermissionResponse]{}, err

	}

	var data []auth_dto.PermissionResponse

	for _, permission := range result.Data {
		data = append(data, auth_dto.ToPermissionResponse(permission))
	}

	return helper.PaginatedResponse[auth_dto.PermissionResponse]{
		Data: data,
		Meta: result.Meta,
	}, nil
}

func (s *permissionService) Show(ctx context.Context, id uint) (*auth_dto.PermissionResponse, error) {
	var permission auth_model.Permission
	if err := s.db.First(&permission, id).Error; err != nil {
		return nil, err
	}

	response := auth_dto.ToPermissionResponse(permission)
	return &response, nil
}

func (s *permissionService) Update(ctx context.Context, id uint, data auth_dto.UpdatePermission) (*auth_dto.PermissionResponse, error) {
	var permission auth_model.Permission
	if err := s.db.First(&permission, id).Error; err != nil {
		return nil, err
	}

	permission.Name = data.Name
	permission.IsActive = *data.IsActive

	if err := s.db.Save(&permission).Error; err != nil {
		return nil, err
	}

	response := auth_dto.ToPermissionResponse(permission)
	return &response, nil
}
