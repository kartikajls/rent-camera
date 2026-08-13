package handler

import "github.com/labstack/echo/v4"

type RentalOrderDetailHandler interface {
	CreateDetail(c echo.Context) error
	GetDetailByID(c echo.Context) error
	GetDetailsByOrderID(c echo.Context) error

	// Admin
	GetAllDetails(c echo.Context) error
	UpdateDetail(c echo.Context) error
	DeleteDetail(c echo.Context) error
}
