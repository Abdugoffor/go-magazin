package cart_model

import "time"

type Cart struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"size:100"`
	Description string `gorm:"size:255"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
