package auth_dto

import (
	"category-crud/helper"
	auth_model "category-crud/module/auth/model"
)

type CreateUser struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Roles    []uint `json:"roles" validate:"required"`
	Password string `json:"password" validate:"required"`
	IsActive bool   `json:"is_active" validate:"required"`
}

type UpdateUser struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Roles    []uint `json:"roles" validate:"required"`
	Password string `json:"password"`
	IsActive bool   `json:"is_active"`
}

type LoginUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserResponse struct {
	ID        uint               `json:"id"`
	Name      string             `json:"name"`
	Email     string             `json:"email"`
	Roles     []RoleUserResponse `json:"roles"`
	IsActive  bool               `json:"is_active"`
	CreatedAt string             `json:"created_at"`
	UpdatedAt string             `json:"updated_at"`
}

func ToUserResponse(user auth_model.User) UserResponse {
	var roles []RoleUserResponse

	for _, r := range user.Roles {

		roles = append(roles, RoleUserResponse{
			ID:       r.ID,
			Name:     r.Name,
			IsActive: r.IsActive,
		})
	}

	return UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Roles:     roles,
		IsActive:  user.IsActive,
		CreatedAt: helper.FormatDate(user.CreatedAt),
		UpdatedAt: helper.FormatDate(user.UpdatedAt),
	}
}
