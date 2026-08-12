package usecase

import (
	"p2-ip-kartikajls/internal/dto"
)

type TopUpUsecase interface {
	//User
	Create(userID int64, request dto.CreateTopUpRequest) (*dto.TopUpResponse, error)
	GetByID(id int64) (*dto.TopUpResponse, error)
	GetByUserID(userID int64) ([]dto.TopUpResponse, error)

	//Admin
	GetAll() ([]dto.TopUpResponse, error)
	UpdateStatus(id int64, status string) error
	Approve(id int64) error
	Reject(id int64) error
}
