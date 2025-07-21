package order_model

import "time"

type Order struct {
	ID         int          `json:"id" gorm:"primaryKey"`
	UserID     int          `json:"user_id"`
	Summ       int          `json:"summ"`
	OrderItems []OrderItems `json:"order_items" gorm:"foreignKey:OrderID"`
	Status     string       `json:"status"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (*Order) TableName() string {
	return "orders"
}
