package auth_dto

import (
	"category-crud/helper"
	auth_model "category-crud/module/auth/model"
)

type CreateRole struct {
	Name        string `json:"name" validate:"required"`
	Permissions []uint `json:"permissions" validate:"required,min=1"`
	IsActive    *bool  `json:"is_active" validate:"required"`
}

type UpdateRole struct {
	Name        string `json:"name" validate:"required"`
	Permissions []uint `json:"permissions" validate:"required,min=1"`
	IsActive    *bool  `json:"is_active" validate:"required"`
}

type RoleResponse struct {
	ID          uint                 `json:"id"`
	Name        string               `json:"name"`
	Permissions []PermissionResponse `json:"permissions"`
	IsActive    bool                 `json:"is_active"`
	CreatedAt   string               `json:"created_at"`
	UpdatedAt   string               `json:"updated_at"`
}

type RoleUserResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}

func ToResponse(role auth_model.Role) RoleResponse {
	var perms []PermissionResponse
	for _, p := range role.Permissions {
		perms = append(perms, PermissionResponse{
			ID:       p.ID,
			Name:     p.Name,
			Path:     p.Path,
			IsActive: p.IsActive,
		})
	}

	return RoleResponse{
		ID:          role.ID,
		Name:        role.Name,
		Permissions: perms,
		IsActive:    role.IsActive,
		CreatedAt:   helper.FormatDate(role.CreatedAt),
		UpdatedAt:   helper.FormatDate(role.UpdatedAt),
	}
}
