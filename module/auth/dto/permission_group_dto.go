package auth_dto

import (
	"category-crud/helper"
	auth_model "category-crud/module/auth/model"
)

type UpdatePermissionGroup struct {
	Name     string `json:"name" validate:"required"`
	IsActive *bool  `json:"is_active" validate:"required"`
}

type PermissionGroupResponse struct {
	ID          uint                 `json:"id"`
	Name        string               `json:"name"`
	Permissions []PermissionResponse `json:"permissions"`
	IsActive    bool                 `json:"is_active"`
	CreatedAt   string               `json:"created_at"`
	UpdatedAt   string               `json:"updated_at"`
}

func ToPermissionGroupResponse(permission_groups auth_model.PermissionGroup) PermissionGroupResponse {
	var perms []PermissionResponse
	for _, p := range permission_groups.Permissions {
		perms = append(perms, PermissionResponse{
			ID:       p.ID,
			Name:     p.Name,
			Path:     p.Path,
			IsActive: p.IsActive,
		})
	}

	return PermissionGroupResponse{
		ID:          permission_groups.ID,
		Name:        permission_groups.Name,
		Permissions: perms,
		IsActive:    permission_groups.IsActive,
		CreatedAt:   helper.FormatDate(permission_groups.CreatedAt),
		UpdatedAt:   helper.FormatDate(permission_groups.UpdatedAt),
	}
}
