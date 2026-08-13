package dto

type CreateCameraRequest struct {
	Name       string  `json:"name"`
	RentalCost float64 `json:"rental_cost"`
	Category   string  `json:"category"`
}

type UpdateCameraRequest struct {
	Name       string  `json:"name"`
	RentalCost float64 `json:"rental_cost"`
	Category   string  `json:"category"`
	Available  *bool   `json:"available"`
}
