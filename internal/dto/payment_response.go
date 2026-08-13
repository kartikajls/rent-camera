package dto

import "time"

type PaymentResponse struct {
	PaymentID     int64      `json:"payment_id"`
	RentalOrderID int64      `json:"rental_order_id"`
	Amount        float64    `json:"amount"`
	PaymentMethod string     `json:"payment_method"`
	PaymentStatus string     `json:"payment_status"`
	PaymentDate   *time.Time `json:"payment_date"`
	CreatedAt     time.Time  `json:"created_at"`
}
