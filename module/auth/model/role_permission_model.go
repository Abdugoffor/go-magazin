package auth_model

import "time"

type RolePermission struct {
	ID           uint `gorm:"primaryKey" json:"id"`
	RoleID       uint `json:"role_id"`
	PermissionID uint `json:"permission_id"`
	IsActive     bool `gorm:"default:true" json:"is_active"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (RolePermission) TableName() string {
	return "role_permissions"
}
