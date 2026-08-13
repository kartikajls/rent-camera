package dto

type CreatePaymentRequest struct {
	RentalOrderID int64  `json:"rental_order_id" validate:"required"`
	PaymentMethod string `json:"payment_method" validate:"required"`
}

type UpdatePaymentStatusRequest struct {
	PaymentStatus string `json:"payment_status" validate:"required"`
}
