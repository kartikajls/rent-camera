package dto

type CreateRentalOrderRequest struct {
	CameraID int64 `json:"camera_id" validate:"required"`
}

type UpdateRentalOrderStatusRequest struct {
	Status string `json:"status" validate:"required"`
}
