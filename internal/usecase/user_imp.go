package usecase

import (
	"errors"

	"p2-ip-kartikajls/internal/dto"
	"p2-ip-kartikajls/internal/entity"
	"p2-ip-kartikajls/internal/helper"
	"p2-ip-kartikajls/internal/repository"
)

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(
	userRepository repository.UserRepository,
) UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}

func (u *userUsecase) Register(request dto.RegisterRequest) (*dto.RegisterResponse, error) {

	// Validasi username
	if request.Username == "" {
		return nil, errors.New("username is required")
	}

	// Validasi email
	if request.Email == "" {
		return nil, errors.New("email is required")
	}

	// Validasi password
	if request.Password == "" {
		return nil, errors.New("password is required")
	}

	// Cek email sudah terdaftar
	existingUser, err := u.userRepository.GetByEmail(
		request.Email,
	)

	if err == nil && existingUser != nil {
		return nil, errors.New("email already registered")
	}

	// Buat entity user
	user := &entity.User{
		Username:      request.Username,
		Email:         request.Email,
		Password:      request.Password,
		Role:          "user",
		DepositAmount: 0,
	}

	// Simpan user
	err = u.userRepository.Create(user)

	if err != nil {
		return nil, err
	}

	// Response
	return &dto.RegisterResponse{
		UserID:        user.UserID,
		Username:      user.Username,
		Email:         user.Email,
		DepositAmount: user.DepositAmount,
		CreatedAt:     user.CreatedAt,
	}, nil
}

func (u *userUsecase) Login(request dto.LoginRequest) (*dto.LoginResponse, error) {

	user, err := u.userRepository.GetByEmail(
		request.Email,
	)

	if err != nil {
		return nil, errors.New(
			"email or password is incorrect",
		)
	}

	if user.Password != request.Password {
		return nil, errors.New(
			"email or password is incorrect",
		)
	}

	// Generate JWT
	token, err := helper.GenerateJWT(
		user.UserID,
		user.Email,
		user.Role,
	)

	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		UserID:   user.UserID,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
		JWTToken: token,
	}, nil
}

func (u *userUsecase) GetByID(userID int64) (*dto.UserResponse, error) {

	user, err := u.userRepository.GetByID(userID)

	if err != nil {
		return nil, err
	}

	return &dto.UserResponse{
		UserID:        user.UserID,
		Username:      user.Username,
		Email:         user.Email,
		DepositAmount: user.DepositAmount,
		CreatedAt:     user.CreatedAt,
	}, nil
}

func (u *userUsecase) GetAll() ([]dto.UserResponse, error) {

	users, err := u.userRepository.GetAll()

	if err != nil {
		return nil, err
	}

	responses := make(
		[]dto.UserResponse,
		0,
		len(users),
	)

	for _, user := range users {

		responses = append(
			responses,
			dto.UserResponse{
				UserID:        user.UserID,
				Username:      user.Username,
				Email:         user.Email,
				Role:          user.Role,
				CreatedAt:     user.CreatedAt,
				DepositAmount: user.DepositAmount,
			},
		)
	}

	return responses, nil
}
