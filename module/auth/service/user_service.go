package auth_service

import (
	"category-crud/helper"
	auth_dto "category-crud/module/auth/dto"
	auth_model "category-crud/module/auth/model"
	"context"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserService interface {
	All(ctx echo.Context) (helper.PaginatedResponse[auth_dto.UserResponse], error)
	Show(ctx context.Context, id uint) (*auth_dto.UserResponse, error)
	Create(ctx context.Context, data auth_dto.CreateUser) (*auth_dto.UserResponse, error)
	Register(ctx context.Context, data auth_dto.RegisterUser) (*auth_dto.UserResponse, error)
	Update(ctx context.Context, id uint, data auth_dto.UpdateUser) (*auth_dto.UserResponse, error)
	Delete(ctx context.Context, id uint) error
}

type userService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{
		db: db,
	}
}

func (s *userService) All(ctx echo.Context) (helper.PaginatedResponse[auth_dto.UserResponse], error) {
	var users []auth_model.User

	result, err := helper.Paginate(ctx, s.db.Preload("Roles"), &users, 10)
	if err != nil {
		return helper.PaginatedResponse[auth_dto.UserResponse]{}, err
	}

	var data []auth_dto.UserResponse
	for _, user := range result.Data {
		data = append(data, auth_dto.ToUserResponse(user))
	}

	return helper.PaginatedResponse[auth_dto.UserResponse]{
		Data: data,
		Meta: result.Meta,
	}, nil
}

func (s *userService) Show(ctx context.Context, id uint) (*auth_dto.UserResponse, error) {
	var user auth_model.User
	if err := s.db.Preload("Roles").First(&user, id).Error; err != nil {
		return nil, err
	}

	response := auth_dto.ToUserResponse(user)
	return &response, nil
}

func (s *userService) Create(ctx context.Context, req auth_dto.CreateUser) (*auth_dto.UserResponse, error) {
	hashedPassword, err := helper.HashPassword(req.Password)

	if err != nil {
		return nil, err
	}

	user := auth_model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
		IsActive: req.IsActive,
	}

	if err := s.db.Create(&user).Error; err != nil {
		return nil, err

	}

	var userRoles []auth_model.Role

	if len(req.Roles) > 0 {
		if err := s.db.Where("id IN ?", req.Roles).Find(&userRoles).Error; err != nil {
			return nil, err
		}

		for _, r := range userRoles {
			ur := auth_model.RoleUser{
				UserID:   user.ID,
				RoleID:   r.ID,
				IsActive: true,
			}
			if err := s.db.Create(&ur).Error; err != nil {
				return nil, err
			}
		}
	}

	if err := s.db.Preload("Roles").First(&user, user.ID).Error; err != nil {
		return nil, err
	}

	response := auth_dto.ToUserResponse(user)
	return &response, nil
}

func (s *userService) Register(ctx context.Context, req auth_dto.RegisterUser) (*auth_dto.UserResponse, error) {
	hashedPassword, err := helper.HashPassword(req.Password)

	if err != nil {
		return nil, err
	}

	user := auth_model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
		IsActive: req.IsActive,
	}

	if err := s.db.Create(&user).Error; err != nil {
		return nil, err

	}

	var role auth_model.Role

	if err := s.db.Where("id = ?", 3).First(&role).Error; err != nil {
		return nil, err
	}

	ur := auth_model.RoleUser{
		UserID:   user.ID,
		RoleID:   role.ID,
		IsActive: true,
	}

	if err := s.db.Create(&ur).Error; err != nil {
		return nil, err
	}

	if err := s.db.Preload("Roles").First(&user, user.ID).Error; err != nil {
		return nil, err
	}

	token, err := helper.GenerateToken(user.ID)

	if err != nil {
		return nil, err
	}

	response := auth_dto.ToUserResponse(user)

	response.Token = token

	return &response, nil
}

func (s *userService) Update(ctx context.Context, id uint, data auth_dto.UpdateUser) (*auth_dto.UserResponse, error) {
	var user auth_model.User
	if err := s.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	if data.Password != "" {
		hashedPassword, err := helper.HashPassword(data.Password)
		if err != nil {
			return nil, err
		}
		user.Password = hashedPassword
	}

	user.Name = data.Name
	user.Email = data.Email
	user.IsActive = data.IsActive

	if err := s.db.Save(&user).Error; err != nil {
		return nil, err
	}

	if len(data.Roles) > 0 {
		if err := s.db.Where("user_id = ?", user.ID).Delete(&auth_model.RoleUser{}).Error; err != nil {
			return nil, err
		}

		for _, roleID := range data.Roles {
			ur := auth_model.RoleUser{
				UserID:   user.ID,
				RoleID:   roleID,
				IsActive: true,
			}
			if err := s.db.Create(&ur).Error; err != nil {
				return nil, err
			}
		}
	}

	if err := s.db.Preload("Roles").First(&user, user.ID).Error; err != nil {
		return nil, err
	}

	response := auth_dto.ToUserResponse(user)
	return &response, nil

}

func (s *userService) Delete(ctx context.Context, id uint) error {
	var user auth_model.User
	if err := s.db.First(&user, id).Error; err != nil {
		return err
	}

	if err := s.db.Where("user_id = ?", user.ID).Delete(&auth_model.RoleUser{}).Error; err != nil {
		return err
	}

	if err := s.db.Delete(&user).Error; err != nil {
		return err
	}

	return nil
}
