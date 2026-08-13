package dto

type CreateRentalOrderRequest struct {
	TotalAmount float64 `json:"total_amount" validate:"required,min=0"`
}

type UpdateRentalOrderStatusRequest struct {
	Status string `json:"status" validate:"required"`
}
