package usecase

import (
	"p2-ip-kartikajls/internal/dto"
)

type RentalOrderUsecase interface {
	CreateOrder(userID int64, req dto.CreateRentalOrderRequest) (*dto.RentalOrderResponse, error)
	GetOrderByID(orderID int64) (*dto.RentalOrderResponse, error)
	GetOrdersByUserID(userID int64) ([]dto.RentalOrderResponse, error)
	GetAllOrders() ([]dto.RentalOrderResponse, error)
	UpdateOrderStatus(orderID int64, status string) error
	DeleteOrder(orderID int64) error
}
