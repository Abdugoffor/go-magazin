package auth_model

import "time"

type Permission struct {
	ID        uint             `gorm:"primaryKey" json:"id"`
	Name      string           `gorm:"not null" json:"name"`
	Path      string           `gorm:"unique;not null" json:"path"`
	GroupID   *uint            `json:"group_id"`
	Group     *PermissionGroup `gorm:"foreignKey:GroupID" json:"group,omitempty"`
	IsActive  bool             `gorm:"default:true" json:"is_active"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
}

func (Permission) TableName() string {
	return "permissions"
}
