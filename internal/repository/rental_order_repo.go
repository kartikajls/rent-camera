package repository

import (
	"p2-ip-kartikajls/internal/dto"
)

type RentalOrderRepository interface {
	CreateOrder(userID int64, req dto.CreateRentalOrderRequest) (*dto.RentalOrderResponse, error)
	GetOrderByID(orderID int64) (*dto.RentalOrderResponse, error)
	GetOrdersByUserID(userID int64) ([]dto.RentalOrderResponse, error)
	GetAllOrders() ([]dto.RentalOrderResponse, error)
	UpdateOrderStatus(orderID int64, status string) error
	DeleteOrder(orderID int64) error
}
