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

// CreateOrder godoc
// @Summary     Create Rental Order
// @Description User membuat order rental dengan memilih camera
// @Tags        User - Order
// @Accept      application/json
// @Produce     application/json
// @Security    BearerAuth
// @Param       request body dto.CreateRentalOrderRequest true "Rental Order Request"
// @Success     201 {object} helper.Response
// @Failure     400 {object} helper.Response
// @Failure     401 {object} helper.Response
// @Router      /users/orders [post]
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

// GetMyOrders godoc
// @Summary Get My Orders
// @Description Mendapatkan seluruh order milik user yang sedang login
// @Tags User - Order
// @Produce json
// @Security BearerAuth
// @Success 200 {object} helper.Response
// @Failure 401 {object} helper.Response
// @Router /users/orders [get]
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

// GetOrderByID godoc
// @Summary     Get Order By ID
// @Description User melihat detail order
// @Tags        User - Order
// @Accept      application/json
// @Produce     application/json
// @Security    BearerAuth
// @Param       id path int64 true "Rental Order ID"
// @Success     200 {object} helper.Response
// @Failure     401 {object} helper.Response
// @Failure     404 {object} helper.Response
// @Router      /users/orders/{id} [get]
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

// GetAllOrders godoc
// @Summary Get All Orders
// @Description Admin melihat seluruh rental order
// @Tags Admin - Order Approve
// @Produce json
// @Security BearerAuth
// @Success 200 {object} helper.Response
// @Failure 401 {object} helper.Response
// @Failure 403 {object} helper.Response
// @Router /admin/orders [get]
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

// UpdateOrderStatus godoc
// @Summary     Approve Rental Order
// @Description Admin menyetujui atau mengubah status rental order user
// @Tags        Admin - Order Approve
// @Accept      application/json
// @Produce     application/json
// @Security    BearerAuth
// @Param       id path int64 true "Rental Order ID"
// @Param       request body dto.UpdateRentalOrderStatusRequest true "Order Status Request"
// @Success     200 {object} helper.Response
// @Failure     400 {object} helper.Response
// @Failure     401 {object} helper.Response
// @Failure     403 {object} helper.Response
// @Failure     404 {object} helper.Response
// @Router      /admin/orders/{id}/status [put]
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
