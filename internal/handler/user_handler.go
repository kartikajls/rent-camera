package handler

import "github.com/labstack/echo/v4"

type UserHandler interface {
	Register(c echo.Context) error
	Login(c echo.Context) error
	GetByID(c echo.Context) error
	GetAll(c echo.Context) error
}
