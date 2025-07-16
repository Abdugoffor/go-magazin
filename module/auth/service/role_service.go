package auth_service

import (
	"category-crud/helper"
	auth_dto "category-crud/module/auth/dto"
	auth_model "category-crud/module/auth/model"
	"context"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type RoleService interface {
	All(ctx echo.Context) (helper.PaginatedResponse[auth_dto.RoleResponse], error)
	Show(ctx context.Context, id uint) (*auth_dto.RoleResponse, error)
	Create(ctx context.Context, data auth_dto.CreateRole) (*auth_dto.RoleResponse, error)
	Update(ctx context.Context, id uint, data auth_dto.UpdateRole) (*auth_dto.RoleResponse, error)
	Delete(ctx context.Context, id uint) error
}

type roleService struct {
	db *gorm.DB
}

func NewRoleService(db *gorm.DB) RoleService {
	return &roleService{
		db: db,
	}
}

func (s *roleService) All(ctx echo.Context) (helper.PaginatedResponse[auth_dto.RoleResponse], error) {
	var roles []auth_model.Role

	result, err := helper.Paginate(ctx, s.db.Preload("Permissions"), &roles, 10)
	if err != nil {
		return helper.PaginatedResponse[auth_dto.RoleResponse]{}, err
	}

	var data []auth_dto.RoleResponse
	for _, role := range result.Data {
		data = append(data, auth_dto.ToResponse(role))
	}

	return helper.PaginatedResponse[auth_dto.RoleResponse]{
		Data: data,
		Meta: result.Meta,
	}, nil
}

func (s *roleService) Show(ctx context.Context, id uint) (*auth_dto.RoleResponse, error) {
	var role auth_model.Role
	if err := s.db.Preload("Permissions").First(&role, id).Error; err != nil {
		return nil, err
	}

	response := auth_dto.ToResponse(role)
	return &response, nil
}

func (s *roleService) Create(ctx context.Context, req auth_dto.CreateRole) (*auth_dto.RoleResponse, error) {

	role := auth_model.Role{
		Name:     req.Name,
		IsActive: *req.IsActive,
	}

	if err := s.db.Create(&role).Error; err != nil {
		return nil, err
	}

	var permissions []auth_model.Permission

	if len(req.Permissions) > 0 {
		if err := s.db.Where("id IN ?", req.Permissions).Find(&permissions).Error; err != nil {
			return nil, err
		}

		for _, p := range permissions {
			rp := auth_model.RolePermission{
				RoleID:       role.ID,
				PermissionID: p.ID,
				IsActive:     true,
			}
			if err := s.db.Create(&rp).Error; err != nil {
				return nil, err
			}
		}

		role.Permissions = permissions
	}

	resp := auth_dto.ToResponse(role)

	return &resp, nil
}

func (s *roleService) Update(ctx context.Context, id uint, req auth_dto.UpdateRole) (*auth_dto.RoleResponse, error) {
	var role auth_model.Role

	if err := s.db.First(&role, id).Error; err != nil {
		return nil, err
	}

	role.Name = req.Name
	role.IsActive = *req.IsActive

	if err := s.db.Save(&role).Error; err != nil {
		return nil, err
	}

	if err := s.db.Where("role_id = ?", role.ID).Delete(&auth_model.RolePermission{}).Error; err != nil {
		return nil, err
	}

	var permissions []auth_model.Permission

	if len(req.Permissions) > 0 {
		if err := s.db.Where("id IN ?", req.Permissions).Find(&permissions).Error; err != nil {
			return nil, err
		}
		for _, p := range permissions {
			rp := auth_model.RolePermission{
				RoleID:       role.ID,
				PermissionID: p.ID,
				IsActive:     true,
			}
			if err := s.db.Create(&rp).Error; err != nil {
				return nil, err
			}
		}

		role.Permissions = permissions
	}

	resp := auth_dto.ToResponse(role)

	return &resp, nil

}

func (s *roleService) Delete(ctx context.Context, id uint) error {
	var role auth_model.Role

	// Role mavjudligini tekshiramiz
	if err := s.db.First(&role, id).Error; err != nil {
		return err
	}

	// Avval role_permissions jadvalidan o‘chiramiz
	if err := s.db.Where("role_id = ?", role.ID).Delete(&auth_model.RolePermission{}).Error; err != nil {
		return err
	}

	// So‘ngra o‘sha Role ni o‘chiramiz
	if err := s.db.Delete(&role).Error; err != nil {
		return err
	}

	return nil
}
