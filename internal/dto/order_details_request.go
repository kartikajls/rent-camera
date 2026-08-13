package dto

type CreateRentalOrderDetailRequest struct {
	CameraID   int64   `json:"camera_id" validate:"required"`
	RentalCost float64 `json:"rental_cost" validate:"required,min=0"`
	RentalDays int     `json:"rental_days" validate:"required,min=1"`
}

type UpdateRentalOrderDetailRequest struct {
	CameraID   int64   `json:"camera_id" validate:"required"`
	RentalCost float64 `json:"rental_cost" validate:"required,min=0"`
	RentalDays int     `json:"rental_days" validate:"required,min=1"`
}
