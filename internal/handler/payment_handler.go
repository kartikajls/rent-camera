package handler

import (
	"github.com/labstack/echo/v4"
)

type PaymentHandler interface {
	// USER
	CreatePayment(c echo.Context) error
	GetMyPayments(c echo.Context) error
	GetMyPaymentByID(c echo.Context) error
	// ADMIN
	GetAllPayments(c echo.Context) error
	UpdatePaymentStatus(c echo.Context) error
}
