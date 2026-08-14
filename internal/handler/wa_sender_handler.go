package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"p2-ip-kartikajls/internal/dto"
	"p2-ip-kartikajls/internal/helper"
	"p2-ip-kartikajls/internal/usecase"
)

type WasenderHandler struct {
	wasenderUsecase usecase.WasenderUsecase
}

func NewWasenderHandler(
	wasenderUsecase usecase.WasenderUsecase,
) *WasenderHandler {
	return &WasenderHandler{
		wasenderUsecase: wasenderUsecase,
	}
}

func (h *WasenderHandler) SendMessage(c echo.Context) error {

	var req dto.SendWARequest

	// Bind request
	if err := c.Bind(&req); err != nil {
		return helper.ResponseError(
			c,
			http.StatusBadRequest,
			"Invalid request body",
			nil,
		)
	}

	// Validasi phone
	if req.Phone == "" {
		return helper.ResponseError(
			c,
			http.StatusBadRequest,
			"Phone is required",
			nil,
		)
	}

	// Validasi message
	if req.Message == "" {
		return helper.ResponseError(
			c,
			http.StatusBadRequest,
			"Message is required",
			nil,
		)
	}

	// Kirim WhatsApp
	result, err := h.wasenderUsecase.SendMessage(req)
	if err != nil {
		return helper.ResponseError(
			c,
			http.StatusBadGateway,
			"Failed to send WhatsApp message",
			err.Error(),
		)
	}

	// Response berhasil
	return helper.ResponseSuccess(
		c,
		http.StatusOK,
		"WhatsApp message sent successfully",
		result,
	)
}
