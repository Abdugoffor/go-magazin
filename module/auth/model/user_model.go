package auth_model

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Name      string `json:"name"`
	Email     string `gorm:"unique" json:"email"`
	Password  string `json:"-"`
	Roles     []Role `gorm:"many2many:role_users;"`
	IsActive  bool   `gorm:"default:true" json:"is_active"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (User) TableName() string {
	return "users"
}
