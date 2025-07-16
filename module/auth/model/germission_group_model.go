package auth_model

import "time"

type PermissionGroup struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"unique" json:"name"`
	IsActive  bool   `gorm:"default:true" json:"is_active"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (PermissionGroup) TableName() string {
	return "permission_groups"
}
