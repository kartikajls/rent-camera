package dto

import "time"

type RentalOrderResponse struct {
	RentalOrderID int64     `json:"rental_order_id"`
	UserID        int64     `json:"user_id"`
	OrderDate     time.Time `json:"order_date"`
	TotalAmount   float64   `json:"total_amount"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
