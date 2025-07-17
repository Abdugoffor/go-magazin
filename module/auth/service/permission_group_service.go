package auth_service

import (
	"category-crud/helper"
	auth_dto "category-crud/module/auth/dto"
	auth_model "category-crud/module/auth/model"
	"context"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type PermissionGroupService interface {
	All(ctx echo.Context) (helper.PaginatedResponse[auth_dto.PermissionGroupResponse], error)
	Show(ctx context.Context, id uint) (*auth_dto.PermissionGroupResponse, error)
	Update(ctx context.Context, id uint, data auth_dto.UpdatePermissionGroup) (*auth_dto.PermissionGroupResponse, error)
}

type permissionGroupService struct {
	db *gorm.DB
}

func NewPermissionGroupService(db *gorm.DB) PermissionGroupService {
	return &permissionGroupService{
		db: db,
	}
}

func (s *permissionGroupService) All(ctx echo.Context) (helper.PaginatedResponse[auth_dto.PermissionGroupResponse], error) {
	var permissionGroups []auth_model.PermissionGroup

	result, err := helper.Paginate(ctx, s.db.Preload("Permissions"), &permissionGroups, 10)
	if err != nil {
		return helper.PaginatedResponse[auth_dto.PermissionGroupResponse]{}, err

	}

	var data []auth_dto.PermissionGroupResponse

	for _, permissionGroup := range result.Data {
		data = append(data, auth_dto.ToPermissionGroupResponse(permissionGroup))
	}

	return helper.PaginatedResponse[auth_dto.PermissionGroupResponse]{
		Data: data,
		Meta: result.Meta,
	}, nil
}

func (s *permissionGroupService) Show(ctx context.Context, id uint) (*auth_dto.PermissionGroupResponse, error) {
	var permissionGroup auth_model.PermissionGroup
	if err := s.db.Preload("Permissions").First(&permissionGroup, id).Error; err != nil {
		return nil, err
	}

	response := auth_dto.ToPermissionGroupResponse(permissionGroup)
	return &response, nil
}

func (s *permissionGroupService) Update(ctx context.Context, id uint, data auth_dto.UpdatePermissionGroup) (*auth_dto.PermissionGroupResponse, error) {
	var permissionGroup auth_model.PermissionGroup
	if err := s.db.First(&permissionGroup, id).Error; err != nil {
		return nil, err
	}

	permissionGroup.Name = data.Name
	permissionGroup.IsActive = *data.IsActive

	if err := s.db.Save(&permissionGroup).Error; err != nil {
		return nil, err
	}

	if err := s.db.Preload("Permissions").First(&permissionGroup, id).Error; err != nil {
		return nil, err
	}

	response := auth_dto.ToPermissionGroupResponse(permissionGroup)
	return &response, nil
}
