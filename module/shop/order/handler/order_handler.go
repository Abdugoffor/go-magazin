package order_handler

import "time"

type Order struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"size:100"`
	Description string `gorm:"size:255"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
