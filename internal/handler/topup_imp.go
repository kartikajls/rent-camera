package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"p2-ip-kartikajls/internal/dto"
	"p2-ip-kartikajls/internal/helper"
	"p2-ip-kartikajls/internal/usecase"
)

type topUpHandler struct {
	topUpUsecase usecase.TopUpUsecase
}

func NewTopUpHandler(
	topUpUsecase usecase.TopUpUsecase,
) TopUpHandler {
	return &topUpHandler{
		topUpUsecase: topUpUsecase,
	}
}

// =====================================================
// USER - CREATE TOP UP
// POST /users/topup
// =====================================================

func (h *topUpHandler) Create(c echo.Context) error {

	// Ambil user_id dari JWT Middleware
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

	// Request DTO
	var request dto.CreateTopUpRequest

	if err := c.Bind(&request); err != nil {
		return helper.ResponseError(
			c,
			http.StatusBadRequest,
			"Invalid request body",
			nil,
		)
	}

	// Validasi request
	if request.Amount <= 0 {
		return helper.ResponseError(
			c,
			http.StatusBadRequest,
			"Amount must be greater than 0",
			nil,
		)
	}

	// Panggil usecase
	topUp, err := h.topUpUsecase.Create(
		userID,
		request,
	)

	if err != nil {
		return helper.ResponseError(
			c,
			http.StatusBadRequest,
			err.Error(),
			nil,
		)
	}

	return helper.ResponseSuccess(
		c,
		http.StatusCreated,
		"Top up created successfully",
		topUp,
	)
}

// =====================================================
// USER - GET OWN TOP UPS
// GET /users/topup
// =====================================================

func (h *topUpHandler) GetByUserID(c echo.Context) error {

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

	topUps, err := h.topUpUsecase.GetByUserID(userID)

	if err != nil {
		return helper.ResponseError(
			c,
			http.StatusInternalServerError,
			"Failed to get top ups",
			nil,
		)
	}

	return helper.ResponseSuccess(
		c,
		http.StatusOK,
		"Top ups retrieved successfully",
		topUps,
	)
}

// =====================================================
// USER / ADMIN - GET TOP UP BY ID
// GET /users/topup/:id
// GET /admin/topups/:id
// =====================================================

func (h *topUpHandler) GetByID(c echo.Context) error {

	idParam := c.Param("id")

	id, err := strconv.ParseInt(
		idParam,
		10,
		64,
	)

	if err != nil {
		return helper.ResponseError(
			c,
			http.StatusBadRequest,
			"Invalid top up ID",
			nil,
		)
	}

	topUp, err := h.topUpUsecase.GetByID(id)

	if err != nil {
		return helper.ResponseError(
			c,
			http.StatusNotFound,
			"Top up not found",
			nil,
		)
	}

	return helper.ResponseSuccess(
		c,
		http.StatusOK,
		"Top up retrieved successfully",
		topUp,
	)
}

// =====================================================
// ADMIN - GET ALL TOP UPS
// GET /admin/topups
// =====================================================

func (h *topUpHandler) GetAll(c echo.Context) error {

	topUps, err := h.topUpUsecase.GetAll()

	if err != nil {
		return helper.ResponseError(
			c,
			http.StatusInternalServerError,
			"Failed to get top ups",
			nil,
		)
	}

	return helper.ResponseSuccess(
		c,
		http.StatusOK,
		"All top ups retrieved successfully",
		topUps,
	)
}

// =====================================================
// ADMIN - APPROVE TOP UP
// PUT /admin/topups/:id/approve
// =====================================================

func (h *topUpHandler) Approve(c echo.Context) error {

	idParam := c.Param("id")

	id, err := strconv.ParseInt(
		idParam,
		10,
		64,
	)

	if err != nil {
		return helper.ResponseError(
			c,
			http.StatusBadRequest,
			"Invalid top up ID",
			nil,
		)
	}

	err = h.topUpUsecase.Approve(id)

	if err != nil {
		return helper.ResponseError(
			c,
			http.StatusBadRequest,
			err.Error(),
			nil,
		)
	}

	return helper.ResponseSuccess(
		c,
		http.StatusOK,
		"Top up approved successfully",
		nil,
	)
}

// =====================================================
// ADMIN - REJECT TOP UP
// PUT /admin/topups/:id/reject
// =====================================================

func (h *topUpHandler) Reject(c echo.Context) error {

	idParam := c.Param("id")

	id, err := strconv.ParseInt(
		idParam,
		10,
		64,
	)

	if err != nil {
		return helper.ResponseError(
			c,
			http.StatusBadRequest,
			"Invalid top up ID",
			nil,
		)
	}

	err = h.topUpUsecase.Reject(id)

	if err != nil {
		return helper.ResponseError(
			c,
			http.StatusBadRequest,
			err.Error(),
			nil,
		)
	}

	return helper.ResponseSuccess(
		c,
		http.StatusOK,
		"Top up rejected successfully",
		nil,
	)
}
