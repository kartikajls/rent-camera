package usecase

import "p2-ip-kartikajls/internal/dto"

type CameraUsecase interface {
	// ADMIN
	Create(request dto.CreateCameraRequest) (*dto.CameraResponse, error)
	Update(cameraID int64, request dto.UpdateCameraRequest) (*dto.CameraResponse, error)
	Delete(cameraID int64) error
	// USER + ADMIN
	GetAll() ([]dto.CameraResponse, error)
	GetByID(cameraID int64) (*dto.CameraResponse, error)
}
