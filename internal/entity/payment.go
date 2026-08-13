package entity

import "time"

type Payment struct {
	PaymentID     int64      `gorm:"primaryKey;column:payment_id"`
	RentalOrderID int64      `gorm:"not null;unique;column:rental_order_id"`
	Amount        float64    `gorm:"type:numeric(15,2);not null;column:amount"`
	PaymentMethod string     `gorm:"type:varchar(50);not null;column:payment_method"`
	PaymentStatus string     `gorm:"type:varchar(30);not null;default:PENDING;column:payment_status"`
	PaymentDate   *time.Time `gorm:"column:payment_date"`
	CreatedAt     time.Time  `gorm:"not null;column:created_at"`
}

func (Payment) TableName() string {
	return "payments"
}
