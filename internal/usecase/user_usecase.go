package usecase

import "p2-ip-kartikajls/internal/dto"

type UserUsecase interface {
	Register(request dto.RegisterRequest) (*dto.UserResponse, error)
	Login(request dto.LoginRequest) (*dto.LoginResponse, error)
	GetByID(userID int64) (*dto.UserResponse, error)
	GetAll() ([]dto.UserResponse, error)
}
