package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"p2-ip-kartikajls/internal/dto"
	"p2-ip-kartikajls/internal/helper"
	"p2-ip-kartikajls/internal/usecase"
)

type rentalOrderHandler struct {
	rentalOrderUsecase usecase.RentalOrderUsecase
}

func NewRentalOrderHandler(
	rentalOrderUsecase usecase.RentalOrderUsecase,
) RentalOrderHandler {
	return &rentalOrderHandler{
		rentalOrderUsecase: rentalOrderUsecase,
	}
}

// USER
func (h *rentalOrderHandler) CreateOrder(c echo.Context) error {

	userIDValue := c.Get("user_id")

	if userIDValue == nil {
		return helper.ResponseError(
			c,
			http.StatusUnauthorized,
			"User ID not found",
			nil,
		)
	}

	userID, ok := userIDValue.(int64)

	if !ok {
		return helper.ResponseError(
			c,
			http.StatusUnauthorized,
			"Invalid user ID",
			nil,
		)
	}

	var req dto.CreateRentalOrderRequest

	if err := c.Bind(&req); err != nil {
		return helper.ResponseError(
			c,
			http.StatusBadRequest,
			"Invalid request body",
			nil,
		)
	}

	result, err := h.rentalOrderUsecase.CreateOrder(
		userID,
		req,
	)

	if err != nil {
		return helper.ResponseError(
			c,
			http.StatusInternalServerError,
			err.Error(),
			nil,
		)
	}

	return helper.ResponseSuccess(
		c,
		http.StatusCreated,
		"Rental order created successfully",
		result,
	)
}

func (h *rentalOrderHandler) GetMyOrders(c echo.Context) error {

	userIDValue := c.Get("user_id")

	if userIDValue == nil {
		return helper.ResponseError(
			c,
			http.StatusUnauthorized,
			"User ID not found",
			nil,
		)
	}

	userID, ok := userIDValue.(int64)

	if !ok {
		return helper.ResponseError(
			c,
			http.StatusUnauthorized,
			"Invalid user ID",
			nil,
		)
	}

	result, err := h.rentalOrderUsecase.GetOrdersByUserID(userID)

	if err != nil {
		return helper.ResponseError(
			c,
			http.StatusInternalServerError,
			err.Error(),
			nil,
		)
	}

	return helper.ResponseSuccess(
		c,
		http.StatusOK,
		"Rental orders retrieved successfully",
		result,
	)
}

func (h *rentalOrderHandler) GetOrderByID(c echo.Context) error {

	idParam := c.Param("id")

	orderID, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		return helper.ResponseError(
			c,
			http.StatusBadRequest,
			"Invalid rental order ID",
			nil,
		)
	}

	result, err := h.rentalOrderUsecase.GetOrderByID(orderID)

	if err != nil {
		return helper.ResponseError(
			c,
			http.StatusNotFound,
			err.Error(),
			nil,
		)
	}

	return helper.ResponseSuccess(
		c,
		http.StatusOK,
		"Rental order retrieved successfully",
		result,
	)
}

// ADMIN
func (h *rentalOrderHandler) GetAllOrders(c echo.Context) error {

	result, err := h.rentalOrderUsecase.GetAllOrders()

	if err != nil {
		return helper.ResponseError(
			c,
			http.StatusInternalServerError,
			err.Error(),
			nil,
		)
	}

	return helper.ResponseSuccess(
		c,
		http.StatusOK,
		"All rental orders retrieved successfully",
		result,
	)
}

func (h *rentalOrderHandler) UpdateOrderStatus(c echo.Context) error {

	idParam := c.Param("id")

	orderID, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		return helper.ResponseError(
			c,
			http.StatusBadRequest,
			"Invalid rental order ID",
			nil,
		)
	}

	var req dto.UpdateRentalOrderStatusRequest

	if err := c.Bind(&req); err != nil {
		return helper.ResponseError(
			c,
			http.StatusBadRequest,
			"Invalid request body",
			nil,
		)
	}

	err = h.rentalOrderUsecase.UpdateOrderStatus(
		orderID,
		req.Status,
	)

	if err != nil {
		return helper.ResponseError(
			c,
			http.StatusInternalServerError,
			err.Error(),
			nil,
		)
	}

	return helper.ResponseSuccess(
		c,
		http.StatusOK,
		"Rental order status updated successfully",
		nil,
	)
}

func (h *rentalOrderHandler) DeleteOrder(c echo.Context) error {

	idParam := c.Param("id")

	orderID, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		return helper.ResponseError(
			c,
			http.StatusBadRequest,
			"Invalid rental order ID",
			nil,
		)
	}

	err = h.rentalOrderUsecase.DeleteOrder(orderID)

	if err != nil {
		return helper.ResponseError(
			c,
			http.StatusNotFound,
			err.Error(),
			nil,
		)
	}

	return helper.ResponseSuccess(
		c,
		http.StatusOK,
		"Rental order deleted successfully",
		nil,
	)
}
