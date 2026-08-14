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

// CreatePayment godoc
// @Summary     Create Payment
// @Description User melakukan pembayaran untuk rental order yang telah disetujui
// @Tags        User - Payment
// @Accept      application/json
// @Produce     application/json
// @Security    BearerAuth
// @Param       request body dto.CreatePaymentRequest true "Payment Request"
// @Success     201 {object} helper.Response
// @Failure     400 {object} helper.Response
// @Failure     401 {object} helper.Response
// @Router      /users/payments [post]
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

// GetMyPayments godoc
// @Summary Get My Payments
// @Description Mendapatkan seluruh pembayaran milik user
// @Tags User - Payment
// @Produce json
// @Security BearerAuth
// @Success 200 {object} helper.Response
// @Failure 401 {object} helper.Response
// @Router /users/payments [get]
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

// GetMyPaymentByID godoc
// @Summary Get My Payment By ID
// @Description Mendapatkan detail pembayaran berdasarkan ID
// @Tags User - Payment
// @Produce json
// @Security BearerAuth
// @Param payment_id path int64 true "Payment ID"
// @Success 200 {object} helper.Response
// @Failure 401 {object} helper.Response
// @Failure 404 {object} helper.Response
// @Router /users/payments/{payment_id} [get]
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

// GetAllPayments godoc
// @Summary Get All Payments
// @Description Admin melihat seluruh pembayaran
// @Tags Admin - Payment Approve
// @Produce json
// @Security BearerAuth
// @Success 200 {object} helper.Response
// @Failure 401 {object} helper.Response
// @Failure 403 {object} helper.Response
// @Router /admin/payments [get]
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

// UpdatePaymentStatus godoc
// @Summary     Approve Payment
// @Description Admin menyetujui pembayaran rental order
// @Tags        Admin - Payment Approve
// @Accept      application/json
// @Produce     application/json
// @Security    BearerAuth
// @Param       payment_id path int64 true "Payment ID"
// @Param       request body dto.UpdatePaymentStatusRequest true "Payment Status Request"
// @Success     200 {object} helper.Response
// @Failure     400 {object} helper.Response
// @Failure     401 {object} helper.Response
// @Failure     403 {object} helper.Response
// @Failure     404 {object} helper.Response
// @Router      /admin/payments/{payment_id}/status [put]
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
