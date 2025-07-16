package auth_model

import "time"

type Role struct {
	ID          uint         `gorm:"primaryKey" json:"id"`
	Name        string       `gorm:"unique" json:"name"`
	Users       []User       `gorm:"many2many:role_users;"`
	Permissions []Permission `gorm:"many2many:role_permissions;"`
	IsActive    bool         `gorm:"default:true" json:"is_active"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (Role) TableName() string {
	return "roles"
}
