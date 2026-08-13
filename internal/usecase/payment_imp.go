package usecase

import (
	"p2-ip-kartikajls/internal/dto"
	"p2-ip-kartikajls/internal/repository"
)

type paymentUsecase struct {
	repository repository.PaymentRepository
}

func NewPaymentUsecase(
	repository repository.PaymentRepository,
) PaymentUsecase {
	return &paymentUsecase{
		repository: repository,
	}
}

func (u *paymentUsecase) CreatePayment(userID int64, req dto.CreatePaymentRequest) (*dto.PaymentResponse, error) {

	return u.repository.CreatePayment(
		userID,
		req,
	)
}

func (u *paymentUsecase) GetPaymentByID(paymentID int64, userID int64) (*dto.PaymentResponse, error) {

	return u.repository.GetPaymentByID(
		paymentID,
		userID,
	)
}

func (u *paymentUsecase) GetPaymentsByUserID(userID int64) ([]dto.PaymentResponse, error) {

	return u.repository.GetPaymentsByUserID(
		userID,
	)
}

func (u *paymentUsecase) GetAllPayments() ([]dto.PaymentResponse, error) {

	return u.repository.GetAllPayments()
}

func (u *paymentUsecase) UpdatePaymentStatus(paymentID int64, req dto.UpdatePaymentStatusRequest) (*dto.PaymentResponse, error) {

	return u.repository.UpdatePaymentStatus(
		paymentID,
		req,
	)
}
