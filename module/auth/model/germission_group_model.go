package auth_model

import "time"

type PermissionGroup struct {
	ID          uint         `gorm:"primaryKey" json:"id"`
	Name        string       `gorm:"unique;not null" json:"name"`
	IsActive    bool         `gorm:"default:true" json:"is_active"`
	Permissions []Permission `gorm:"foreignKey:GroupID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"permissions,omitempty"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}

func (PermissionGroup) TableName() string {
	return "permission_groups"
}
