package repository

import (
	"p2-ip-kartikajls/internal/dto"
)

type RentalOrderRepository interface {
	CreateOrder(userID int64, cameraID int64, totalAmount float64) (*dto.RentalOrderResponse, error)
	GetOrderByID(orderID int64) (*dto.RentalOrderResponse, error)
	GetOrdersByUserID(userID int64) ([]dto.RentalOrderResponse, error)
	GetAllOrders() ([]dto.RentalOrderResponse, error)
	UpdateOrderStatus(orderID int64, status string) error
	DeleteOrder(orderID int64) error
}
