package repository

import (
	"errors"
	"time"

	"gorm.io/gorm"

	"p2-ip-kartikajls/internal/dto"
	"p2-ip-kartikajls/internal/entity"
)

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(
	db *gorm.DB,
) PaymentRepository {
	return &paymentRepository{
		db: db,
	}
}

func (r *paymentRepository) CreatePayment(userID int64, req dto.CreatePaymentRequest) (*dto.PaymentResponse, error) {

	var order struct {
		RentalOrderID int64
		UserID        int64
		TotalAmount   float64
	}

	err := r.db.
		Table("rental_orders").
		Select(
			"rental_order_id",
			"user_id",
			"total_amount",
		).
		Where(
			"rental_order_id = ?",
			req.RentalOrderID,
		).
		First(&order).Error

	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(
				"rental order not found",
			)
		}

		return nil, err
	}

	// Pastikan order milik user yang login.
	if order.UserID != userID {
		return nil, errors.New(
			"access denied",
		)
	}

	// Karena rental_order_id UNIQUE,
	// satu order hanya boleh mempunyai satu payment.
	var existing entity.Payment

	err = r.db.
		Where(
			"rental_order_id = ?",
			req.RentalOrderID,
		).
		First(&existing).Error

	if err == nil {
		return nil, errors.New(
			"payment for this rental order already exists",
		)
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	payment := entity.Payment{
		RentalOrderID: req.RentalOrderID,
		Amount:        order.TotalAmount,
		PaymentMethod: req.PaymentMethod,
		PaymentStatus: "PENDING",
	}

	if err := r.db.Create(&payment).Error; err != nil {
		return nil, err
	}

	return &dto.PaymentResponse{
		PaymentID:     payment.PaymentID,
		RentalOrderID: payment.RentalOrderID,
		Amount:        payment.Amount,
		PaymentMethod: payment.PaymentMethod,
		PaymentStatus: payment.PaymentStatus,
		PaymentDate:   payment.PaymentDate,
		CreatedAt:     payment.CreatedAt,
	}, nil
}

func (r *paymentRepository) GetPaymentByID(paymentID int64, userID int64) (*dto.PaymentResponse, error) {

	var result dto.PaymentResponse

	err := r.db.
		Table("payments").
		Select(`
			payments.payment_id,
			payments.rental_order_id,
			payments.amount,
			payments.payment_method,
			payments.payment_status,
			payments.payment_date,
			payments.created_at
		`).
		Joins(`
			JOIN rental_orders
			ON rental_orders.rental_order_id =
			   payments.rental_order_id
		`).
		Where(`
			payments.payment_id = ?
			AND rental_orders.user_id = ?
		`,
			paymentID,
			userID,
		).
		Scan(&result).Error

	if err != nil {
		return nil, err
	}

	if result.PaymentID == 0 {
		return nil, errors.New(
			"payment not found",
		)
	}

	return &result, nil
}

func (r *paymentRepository) GetPaymentsByUserID(userID int64) ([]dto.PaymentResponse, error) {

	var payments []dto.PaymentResponse

	err := r.db.
		Table("payments").
		Select(`
			payments.payment_id,
			payments.rental_order_id,
			payments.amount,
			payments.payment_method,
			payments.payment_status,
			payments.payment_date,
			payments.created_at
		`).
		Joins(`
			JOIN rental_orders
			ON rental_orders.rental_order_id =
			   payments.rental_order_id
		`).
		Where(
			"rental_orders.user_id = ?",
			userID,
		).
		Order(
			"payments.payment_id DESC",
		).
		Scan(&payments).Error

	if err != nil {
		return nil, err
	}

	return payments, nil
}

func (r *paymentRepository) GetAllPayments() ([]dto.PaymentResponse, error) {

	var payments []dto.PaymentResponse

	err := r.db.
		Table("payments").
		Select(`
			payment_id,
			rental_order_id,
			amount,
			payment_method,
			payment_status,
			payment_date,
			created_at
		`).
		Order(
			"payment_id DESC",
		).
		Scan(&payments).Error

	if err != nil {
		return nil, err
	}

	return payments, nil
}

func (r *paymentRepository) UpdatePaymentStatus(paymentID int64, req dto.UpdatePaymentStatusRequest) (*dto.PaymentResponse, error) {

	var payment entity.Payment

	err := r.db.
		Where(
			"payment_id = ?",
			paymentID,
		).
		First(&payment).Error

	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(
				"payment not found",
			)
		}

		return nil, err
	}

	payment.PaymentStatus = req.PaymentStatus

	// Jika payment berhasil, isi payment_date.
	if req.PaymentStatus == "PAID" {
		now := time.Now()
		payment.PaymentDate = &now
	}

	if err := r.db.Save(&payment).Error; err != nil {
		return nil, err
	}

	return &dto.PaymentResponse{
		PaymentID:     payment.PaymentID,
		RentalOrderID: payment.RentalOrderID,
		Amount:        payment.Amount,
		PaymentMethod: payment.PaymentMethod,
		PaymentStatus: payment.PaymentStatus,
		PaymentDate:   payment.PaymentDate,
		CreatedAt:     payment.CreatedAt,
	}, nil
}
