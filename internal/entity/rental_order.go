package entity

import "time"

type RentalOrder struct {
	RentalOrderID int64     `gorm:"primaryKey;column:rental_order_id"`
	UserID        int64     `gorm:"not null;column:user_id"`
	OrderDate     time.Time `gorm:"not null;column:order_date"`
	TotalAmount   float64   `gorm:"type:numeric(15,2);not null;default:0;column:total_amount"`
	Status        string    `gorm:"type:varchar(30);not null;default:PENDING;column:status"`
	CreatedAt     time.Time `gorm:"not null;column:created_at"`
	UpdatedAt     time.Time `gorm:"not null;column:updated_at"`
}

func (RentalOrder) TableName() string {
	return "rental_orders"
}
