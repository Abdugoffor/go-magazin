package auth_model

import "time"

type Permission struct {
	ID        uint            `gorm:"primaryKey" json:"id"`
	Name      string          `json:"name"`
	Path      string          `gorm:"unique" json:"path"`
	GroupID   uint            `json:"group_id"`
	Group     PermissionGroup `gorm:"foreignKey:GroupID"`
	IsActive  bool            `gorm:"default:true" json:"is_active"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Permission) TableName() string {
	return "permissions"
}
