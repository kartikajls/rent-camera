package repository

import (
	"errors"

	"gorm.io/gorm"

	"p2-ip-kartikajls/internal/dto"
	"p2-ip-kartikajls/internal/entity"
)

type rentalOrderRepository struct {
	db *gorm.DB
}

func NewRentalOrderRepository(db *gorm.DB) RentalOrderRepository {
	return &rentalOrderRepository{
		db: db,
	}
}

func (r *rentalOrderRepository) CreateOrder(userID int64, cameraID int64, totalAmount float64) (*dto.RentalOrderResponse, error) {

	order := entity.RentalOrder{
		UserID:      userID,
		TotalAmount: totalAmount,
		Status:      "pending",
	}

	if err := r.db.Create(&order).Error; err != nil {
		return nil, err
	}

	return &dto.RentalOrderResponse{
		RentalOrderID: order.RentalOrderID,
		UserID:        order.UserID,
		OrderDate:     order.OrderDate,
		TotalAmount:   order.TotalAmount,
		Status:        order.Status,
		CreatedAt:     order.CreatedAt,
		UpdatedAt:     order.UpdatedAt,
	}, nil
}

func (r *rentalOrderRepository) GetOrderByID(orderID int64) (*dto.RentalOrderResponse, error) {

	var order entity.RentalOrder

	if err := r.db.
		Where("rental_order_id = ?", orderID).
		First(&order).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("rental order not found")
		}

		return nil, err
	}

	return &dto.RentalOrderResponse{
		RentalOrderID: order.RentalOrderID,
		UserID:        order.UserID,
		OrderDate:     order.OrderDate,
		TotalAmount:   order.TotalAmount,
		Status:        order.Status,
		CreatedAt:     order.CreatedAt,
		UpdatedAt:     order.UpdatedAt,
	}, nil
}

func (r *rentalOrderRepository) GetOrdersByUserID(userID int64) ([]dto.RentalOrderResponse, error) {

	var orders []entity.RentalOrder

	if err := r.db.
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&orders).Error; err != nil {

		return nil, err
	}

	responses := make([]dto.RentalOrderResponse, 0, len(orders))

	for _, order := range orders {
		responses = append(responses, dto.RentalOrderResponse{
			RentalOrderID: order.RentalOrderID,
			UserID:        order.UserID,
			OrderDate:     order.OrderDate,
			TotalAmount:   order.TotalAmount,
			Status:        order.Status,
			CreatedAt:     order.CreatedAt,
			UpdatedAt:     order.UpdatedAt,
		})
	}

	return responses, nil
}

func (r *rentalOrderRepository) GetAllOrders() ([]dto.RentalOrderResponse, error) {

	var orders []entity.RentalOrder

	if err := r.db.
		Order("created_at DESC").
		Find(&orders).Error; err != nil {

		return nil, err
	}

	responses := make([]dto.RentalOrderResponse, 0, len(orders))

	for _, order := range orders {
		responses = append(responses, dto.RentalOrderResponse{
			RentalOrderID: order.RentalOrderID,
			UserID:        order.UserID,
			OrderDate:     order.OrderDate,
			TotalAmount:   order.TotalAmount,
			Status:        order.Status,
			CreatedAt:     order.CreatedAt,
			UpdatedAt:     order.UpdatedAt,
		})
	}

	return responses, nil
}

func (r *rentalOrderRepository) UpdateOrderStatus(orderID int64, status string) error {

	result := r.db.
		Model(&entity.RentalOrder{}).
		Where("rental_order_id = ?", orderID).
		Update("status", status)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("rental order not found")
	}

	return nil
}

func (r *rentalOrderRepository) DeleteOrder(orderID int64) error {

	result := r.db.
		Delete(
			&entity.RentalOrder{},
			"rental_order_id = ?",
			orderID,
		)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("rental order not found")
	}

	return nil
}
