package handler

import (
	"github.com/labstack/echo/v4"
)

type CameraHandler interface {
	Create(c echo.Context) error
	GetAll(c echo.Context) error
	GetByID(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}
