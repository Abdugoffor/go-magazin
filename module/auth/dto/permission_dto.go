package auth_dto

import auth_model "category-crud/module/auth/model"

type RoleWithPermissions struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}

type PermissionResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Path     string `json:"path"`
	IsActive bool   `json:"is_active"`
}

type UpdatePermission struct {
	Name     string `json:"name" validate:"required"`
	IsActive *bool  `json:"is_active" validate:"required"`
}

func ToPermissionResponse(permission auth_model.Permission) PermissionResponse {
	return PermissionResponse{
		ID:       permission.ID,
		Name:     permission.Name,
		Path:     permission.Path,
		IsActive: permission.IsActive,
	}
}
