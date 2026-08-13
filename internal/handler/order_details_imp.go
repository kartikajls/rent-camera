package handler

import (
	"net/http"
	"p2-ip-kartikajls/internal/dto"
	"p2-ip-kartikajls/internal/helper"
	"p2-ip-kartikajls/internal/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

type rentalOrderDetailHandler struct {
	usecase usecase.RentalOrderDetailUsecase
}

func NewRentalOrderDetailHandler(
	usecase usecase.RentalOrderDetailUsecase,
) RentalOrderDetailHandler {
	return &rentalOrderDetailHandler{
		usecase: usecase,
	}
}

func (h *rentalOrderDetailHandler) CreateDetail(c echo.Context) error {

	// Ambil user_id dari JWT
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

	orderID, err := strconv.ParseInt(
		c.Param("order_id"),
		10,
		64,
	)

	if err != nil {
		return helper.ResponseError(
			c,
			http.StatusBadRequest,
			"Invalid rental order ID",
			nil,
		)
	}

	var req dto.CreateRentalOrderDetailRequest

	if err := c.Bind(&req); err != nil {
		return helper.ResponseError(
			c,
			http.StatusBadRequest,
			"Invalid request body",
			nil,
		)
	}

	result, err := h.usecase.CreateDetail(
		orderID,
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

	_ = userID

	return helper.ResponseSuccess(
		c,
		http.StatusCreated,
		"Rental order detail created successfully",
		result,
	)
}

func (h *rentalOrderDetailHandler) GetDetailByID(c echo.Context) error {

	userIDValue := c.Get("user_id")

	if userIDValue == nil {
		return helper.ResponseError(
			c,
			http.StatusUnauthorized,
			"User ID not found",
			nil,
		)
	}

	detailID, err := strconv.ParseInt(
		c.Param("detail_id"),
		10,
		64,
	)

	if err != nil {
		return helper.ResponseError(
			c,
			http.StatusBadRequest,
			"Invalid rental detail ID",
			nil,
		)
	}

	result, err := h.usecase.GetDetailByID(detailID)

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
		"Rental order detail retrieved successfully",
		result,
	)
}

func (h *rentalOrderDetailHandler) GetDetailsByOrderID(c echo.Context) error {

	userIDValue := c.Get("user_id")

	if userIDValue == nil {
		return helper.ResponseError(
			c,
			http.StatusUnauthorized,
			"User ID not found",
			nil,
		)
	}

	orderID, err := strconv.ParseInt(
		c.Param("order_id"),
		10,
		64,
	)

	if err != nil {
		return helper.ResponseError(
			c,
			http.StatusBadRequest,
			"Invalid rental order ID",
			nil,
		)
	}

	result, err := h.usecase.GetDetailsByOrderID(orderID)

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
		"Rental order details retrieved successfully",
		result,
	)
}

func (h *rentalOrderDetailHandler) GetAllDetails(c echo.Context) error {

	result, err := h.usecase.GetAllDetails()

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
		"All rental order details retrieved successfully",
		result,
	)
}

func (h *rentalOrderDetailHandler) UpdateDetail(c echo.Context) error {

	detailID, err := strconv.ParseInt(
		c.Param("detail_id"),
		10,
		64,
	)

	if err != nil {
		return helper.ResponseError(
			c,
			http.StatusBadRequest,
			"Invalid rental detail ID",
			nil,
		)
	}

	var req dto.UpdateRentalOrderDetailRequest

	if err := c.Bind(&req); err != nil {
		return helper.ResponseError(
			c,
			http.StatusBadRequest,
			"Invalid request body",
			nil,
		)
	}

	result, err := h.usecase.UpdateDetail(
		detailID,
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
		http.StatusOK,
		"Rental order detail updated successfully",
		result,
	)
}

func (h *rentalOrderDetailHandler) DeleteDetail(c echo.Context) error {

	detailID, err := strconv.ParseInt(
		c.Param("detail_id"),
		10,
		64,
	)

	if err != nil {
		return helper.ResponseError(
			c,
			http.StatusBadRequest,
			"Invalid rental detail ID",
			nil,
		)
	}

	err = h.usecase.DeleteDetail(detailID)

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
		"Rental order detail deleted successfully",
		nil,
	)
}
