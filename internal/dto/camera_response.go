package dto

import "time"

type CameraResponse struct {
	CameraID   int64     `json:"camera_id"`
	Name       string    `json:"name"`
	Available  bool      `json:"available"`
	RentalCost float64   `json:"rental_cost"`
	Category   string    `json:"category"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
