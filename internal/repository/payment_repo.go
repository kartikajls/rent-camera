package repository

import (
	"p2-ip-kartikajls/internal/dto"
)

type PaymentRepository interface {
	CreatePayment(userID int64, req dto.CreatePaymentRequest) (*dto.PaymentResponse, error)
	GetPaymentByID(paymentID int64, userID int64) (*dto.PaymentResponse, error)
	GetPaymentsByUserID(userID int64) ([]dto.PaymentResponse, error)
	GetAllPayments() ([]dto.PaymentResponse, error)
	UpdatePaymentStatus(paymentID int64, req dto.UpdatePaymentStatusRequest) (*dto.PaymentResponse, error)
}
