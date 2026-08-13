package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"p2-ip-kartikajls/internal/dto"
	"p2-ip-kartikajls/internal/helper"
	"p2-ip-kartikajls/internal/usecase"
)

type cameraHandler struct {
	cameraUsecase usecase.CameraUsecase
}

func NewCameraHandler(
	cameraUsecase usecase.CameraUsecase,
) CameraHandler {
	return &cameraHandler{
		cameraUsecase: cameraUsecase,
	}
}

// ADMIN - CREATE CAMERA
func (h *cameraHandler) Create(c echo.Context) error {

	var request dto.CreateCameraRequest

	if err := c.Bind(&request); err != nil {
		return helper.ResponseError(
			c,
			http.StatusBadRequest,
			"invalid request body",
			nil,
		)
	}

	response, err := h.cameraUsecase.Create(request)

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
		"camera created successfully",
		response,
	)
}

// USER + ADMIN - GET ALL
func (h *cameraHandler) GetAll(c echo.Context) error {

	response, err := h.cameraUsecase.GetAll()

	if err != nil {
		return helper.ResponseError(
			c,
			http.StatusInternalServerError,
			"failed to get cameras",
			nil,
		)
	}

	return helper.ResponseSuccess(
		c,
		http.StatusOK,
		"cameras retrieved successfully",
		response,
	)
}

// USER + ADMIN - GET BY ID
func (h *cameraHandler) GetByID(c echo.Context) error {

	idParam := c.Param("id")

	cameraID, err := strconv.ParseInt(
		idParam,
		10,
		64,
	)

	if err != nil {
		return helper.ResponseError(
			c,
			http.StatusBadRequest,
			"invalid camera id",
			nil,
		)
	}

	response, err := h.cameraUsecase.GetByID(cameraID)

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
		"camera retrieved successfully",
		response,
	)
}

// ADMIN - UPDATE
func (h *cameraHandler) Update(c echo.Context) error {

	idParam := c.Param("id")

	cameraID, err := strconv.ParseInt(
		idParam,
		10,
		64,
	)

	if err != nil {
		return helper.ResponseError(
			c,
			http.StatusBadRequest,
			"invalid camera id",
			nil,
		)
	}

	var request dto.UpdateCameraRequest

	if err := c.Bind(&request); err != nil {
		return helper.ResponseError(
			c,
			http.StatusBadRequest,
			"invalid request body",
			nil,
		)
	}

	response, err := h.cameraUsecase.Update(
		cameraID,
		request,
	)

	if err != nil {
		if err.Error() == "camera not found" {
			return helper.ResponseError(
				c,
				http.StatusNotFound,
				err.Error(),
				nil,
			)
		}

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
		"camera updated successfully",
		response,
	)
}

// ADMIN - DELETE
func (h *cameraHandler) Delete(c echo.Context) error {

	idParam := c.Param("id")

	cameraID, err := strconv.ParseInt(
		idParam,
		10,
		64,
	)

	if err != nil {
		return helper.ResponseError(
			c,
			http.StatusBadRequest,
			"invalid camera id",
			nil,
		)
	}

	err = h.cameraUsecase.Delete(cameraID)

	if err != nil {
		if err.Error() == "camera not found" {
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
		http.StatusOK,
		"camera deleted successfully",
		nil,
	)
}
