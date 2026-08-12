package entity

import "time"

type RentalOrder struct {
	RentalOrderID int64     `gorm:"primaryKey;column:rental_order_id"`
	UserID        int64     `gorm:"not null;column:user_id"`
	OrderDate     time.Time `gorm:"not null;column:order_date"`
	TotalAmount   float64   `gorm:"type:numeric(15,2);not null;column:total_amount"`
	Status        string    `gorm:"type:varchar(50);not null"`
	CreatedAt     time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt     time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (RentalOrder) TableName() string {
	return "rental_orders"
}
