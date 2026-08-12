package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"p2-ip-kartikajls/internal/dto"
	"p2-ip-kartikajls/internal/helper"
	"p2-ip-kartikajls/internal/usecase"
)

type userHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(
	userUsecase usecase.UserUsecase,
) UserHandler {
	return &userHandler{
		userUsecase: userUsecase,
	}
}

// Register godoc
// @Summary Register user
// @Description Register a new user
// @Tags Users
// @Accept json
// @Produce json
// @Param request body dto.RegisterRequest true "Register request"
// @Success 201 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Failure 500 {object} helper.Response
// @Router /users/register [post]
func (h *userHandler) Register(c echo.Context) error {

	var request dto.RegisterRequest

	if err := c.Bind(&request); err != nil {
		return helper.ResponseError(
			c,
			http.StatusBadRequest,
			"Invalid request body",
			nil,
		)
	}

	if request.Username == "" ||
		request.Email == "" ||
		request.Password == "" {

		return helper.ResponseError(
			c,
			http.StatusBadRequest,
			"Username, email, and password are required",
			nil,
		)
	}

	user, err := h.userUsecase.Register(request)

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
		"User registered successfully",
		user,
	)
}

// Login godoc
// @Summary Login user
// @Description Login user and generate JWT
// @Tags Users
// @Accept json
// @Produce json
// @Param request body dto.LoginRequest true "Login request"
// @Success 200 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Failure 401 {object} helper.Response
// @Router /users/login [post]
func (h *userHandler) Login(c echo.Context) error {

	var request dto.LoginRequest

	if err := c.Bind(&request); err != nil {
		return helper.ResponseError(
			c,
			http.StatusBadRequest,
			"Invalid request body",
			nil,
		)
	}

	if request.Email == "" || request.Password == "" {
		return helper.ResponseError(
			c,
			http.StatusBadRequest,
			"Email and password are required",
			nil,
		)
	}

	response, err := h.userUsecase.Login(request)

	if err != nil {
		return helper.ResponseError(
			c,
			http.StatusUnauthorized,
			err.Error(),
			nil,
		)
	}

	return helper.ResponseSuccess(
		c,
		http.StatusOK,
		"Login successful",
		response,
	)
}

// GetByID godoc
// @Summary Get user by ID
// @Description Get user information by ID
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "User ID"
// @Success 200 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Failure 401 {object} helper.Response
// @Failure 404 {object} helper.Response
// @Router /users/{id} [get]
func (h *userHandler) GetByID(c echo.Context) error {

	idParam := c.Param("id")

	userID, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		return helper.ResponseError(
			c,
			http.StatusBadRequest,
			"Invalid user ID",
			nil,
		)
	}

	user, err := h.userUsecase.GetByID(userID)

	if err != nil {
		return helper.ResponseError(
			c,
			http.StatusNotFound,
			"User not found",
			nil,
		)
	}

	return helper.ResponseSuccess(
		c,
		http.StatusOK,
		"User retrieved successfully",
		user,
	)
}

func (h *userHandler) GetAll(c echo.Context) error {

	users, err := h.userUsecase.GetAll()

	if err != nil {
		return helper.ResponseError(
			c,
			http.StatusInternalServerError,
			"Failed to get users",
			nil,
		)
	}

	return helper.ResponseSuccess(
		c,
		http.StatusOK,
		"Users retrieved successfully",
		users,
	)
}
