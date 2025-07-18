package auth_model

import "time"

type RoleUser struct {
	ID        uint `gorm:"primaryKey" json:"id"`
	RoleID    uint `json:"role_id"`
	UserID    uint `json:"user_id"`
	IsActive  bool `gorm:"default:true" json:"is_active"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (RoleUser) TableName() string {
	return "role_users"
}
