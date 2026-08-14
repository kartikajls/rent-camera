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

// Create godoc
// @Summary     Create Top Up
// @Description User membuat permintaan top up saldo
// @Tags        User - Top Up
// @Accept      application/json
// @Produce     application/json
// @Security    BearerAuth
// @Param       request body dto.CreateTopUpRequest true "Top Up Request"
// @Success     201 {object} helper.Response
// @Failure     400 {object} helper.Response
// @Failure     401 {object} helper.Response
// @Router      /users/topup [post]
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

// GetByUserID godoc
// @Summary     Get My Top Ups
// @Description User melihat riwayat top up miliknya
// @Tags        User - Top Up
// @Accept      application/json
// @Produce     application/json
// @Security    BearerAuth
// @Success     200 {object} helper.Response
// @Failure     401 {object} helper.Response
// @Router      /users/topup [get]
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

// USER / ADMIN - GET TOP UP BY ID

// GetByID godoc
// @Summary Get Top Up By ID
// @Description Mendapatkan detail top up berdasarkan ID
// @Tags User - Top Up
// @Produce json
// @Security BearerAuth
// @Param id path int64 true "Top Up ID"
// @Success 200 {object} helper.Response
// @Failure 401 {object} helper.Response
// @Failure 404 {object} helper.Response
// @Router /users/topup/{id} [get]
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

// ADMIN - GET ALL TOP UPS
// GET /admin/topups

// GetAll godoc
// @Summary Get All Top Ups
// @Description Admin melihat seluruh permintaan top up
// @Tags Admin - Top Up Approve
// @Produce json
// @Security BearerAuth
// @Success 200 {object} helper.Response
// @Failure 401 {object} helper.Response
// @Failure 403 {object} helper.Response
// @Router /admin/topups [get]
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

// ADMIN - APPROVE TOP UP

// Approve godoc
// @Summary     Approve Top Up
// @Description Admin menyetujui permintaan top up user
// @Tags        Admin - Top Up Approve
// @Accept      application/json
// @Produce     application/json
// @Security    BearerAuth
// @Param       id path int64 true "Top Up ID"
// @Success     200 {object} helper.Response
// @Failure     400 {object} helper.Response
// @Failure     401 {object} helper.Response
// @Failure     403 {object} helper.Response
// @Failure     404 {object} helper.Response
// @Router      /admin/topups/{id}/approve [put]
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

// Reject godoc
// @Summary      Reject Top Up
// @Description Admin menolak permintaan top up user
// @Tags         Admin - Top Up Reject
// @Accept       application/json
// @Produce      application/json
// @Security     BearerAuth
// @Param        id path int64 true "Top Up ID"
// @Success      200 {object} helper.Response
// @Failure      400 {object} helper.Response
// @Failure      401 {object} helper.Response
// @Failure      403 {object} helper.Response
// @Failure      404 {object} helper.Response
// @Router       /admin/topups/{id}/reject [put]
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
