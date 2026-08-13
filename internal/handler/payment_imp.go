package handler

import (
	"net/http"
	"p2-ip-kartikajls/internal/dto"
	"p2-ip-kartikajls/internal/helper"
	"p2-ip-kartikajls/internal/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

type paymentHandler struct {
	usecase usecase.PaymentUsecase
}

func NewPaymentHandler(
	usecase usecase.PaymentUsecase,
) PaymentHandler {
	return &paymentHandler{
		usecase: usecase,
	}
}

func (h *paymentHandler) CreatePayment(c echo.Context) error {

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

	var req dto.CreatePaymentRequest

	if err := c.Bind(&req); err != nil {
		return helper.ResponseError(
			c,
			http.StatusBadRequest,
			"Invalid request body",
			nil,
		)
	}

	result, err := h.usecase.CreatePayment(
		userID,
		req,
	)

	if err != nil {

		if err.Error() == "access denied" {
			return helper.ResponseError(
				c,
				http.StatusForbidden,
				err.Error(),
				nil,
			)
		}

		if err.Error() ==
			"payment for this rental order already exists" {
			return helper.ResponseError(
				c,
				http.StatusConflict,
				err.Error(),
				nil,
			)
		}

		if err.Error() ==
			"rental order not found" {
			return helper.ResponseError(
				c,
				http.StatusNotFound,
				err.Error(),
				nil,
			)
		}

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
		"Payment created successfully",
		result,
	)
}

func (h *paymentHandler) GetMyPayments(c echo.Context) error {

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

	result, err := h.usecase.GetPaymentsByUserID(
		userID,
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
		"Payments retrieved successfully",
		result,
	)
}

func (h *paymentHandler) GetMyPaymentByID(c echo.Context) error {

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

	paymentID, err := strconv.ParseInt(
		c.Param("payment_id"),
		10,
		64,
	)

	if err != nil {
		return helper.ResponseError(
			c,
			http.StatusBadRequest,
			"Invalid payment ID",
			nil,
		)
	}

	result, err := h.usecase.GetPaymentByID(
		paymentID,
		userID,
	)

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
		"Payment retrieved successfully",
		result,
	)
}

func (h *paymentHandler) GetAllPayments(c echo.Context) error {

	result, err := h.usecase.GetAllPayments()

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
		"All payments retrieved successfully",
		result,
	)
}

func (h *paymentHandler) UpdatePaymentStatus(c echo.Context) error {

	paymentID, err := strconv.ParseInt(
		c.Param("payment_id"),
		10,
		64,
	)

	if err != nil {
		return helper.ResponseError(
			c,
			http.StatusBadRequest,
			"Invalid payment ID",
			nil,
		)
	}

	var req dto.UpdatePaymentStatusRequest

	if err := c.Bind(&req); err != nil {
		return helper.ResponseError(
			c,
			http.StatusBadRequest,
			"Invalid request body",
			nil,
		)
	}

	result, err := h.usecase.UpdatePaymentStatus(
		paymentID,
		req,
	)

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
		"Payment status updated successfully",
		result,
	)
}
