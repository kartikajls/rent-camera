package dto

import "time"

type RentalOrderDetailResponse struct {
	RentalDetailID int64     `json:"rental_detail_id"`
	RentalOrderID  int64     `json:"rental_order_id"`
	CameraID       int64     `json:"camera_id"`
	RentalCost     float64   `json:"rental_cost"`
	RentalDays     int       `json:"rental_days"`
	Subtotal       float64   `json:"subtotal"`
	CreatedAt      time.Time `json:"created_at"`
}
