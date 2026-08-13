package handler

import "github.com/labstack/echo/v4"

type RentalOrderHandler interface {
	CreateOrder(c echo.Context) error
	GetMyOrders(c echo.Context) error
	GetOrderByID(c echo.Context) error
	GetAllOrders(c echo.Context) error
	UpdateOrderStatus(c echo.Context) error
	DeleteOrder(c echo.Context) error
}
