package entity

import "time"

type Payment struct {
	PaymentID     int64      `gorm:"primaryKey;column:payment_id"`
	RentalOrderID int64      `gorm:"not null;column:rental_order_id"`
	Amount        float64    `gorm:"type:numeric(15,2);not null"`
	PaymentMethod string     `gorm:"type:varchar(50);not null;column:payment_method"`
	PaymentStatus string     `gorm:"type:varchar(50);not null;column:payment_status"`
	PaymentProof  *string    `gorm:"type:text;column:payment_proof"`
	PaymentDate   *time.Time `gorm:"column:payment_date"`
	VerifiedAt    *time.Time `gorm:"column:verified_at"`
	VerifiedBy    *int64     `gorm:"column:verified_by"`
	TransactionID *string    `gorm:"type:varchar(255);column:transaction_id"`
	CreatedAt     time.Time  `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt     time.Time  `gorm:"column:updated_at;autoUpdateTime"`
}

func (Payment) TableName() string {
	return "payments"
}
