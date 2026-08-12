package handler

import "github.com/labstack/echo/v4"

type TopUpHandler interface {
	//user
	Create(c echo.Context) error
	GetByID(c echo.Context) error
	GetByUserID(c echo.Context) error

	//Admin
	GetAll(c echo.Context) error
	Approve(c echo.Context) error
	Reject(c echo.Context) error
}
