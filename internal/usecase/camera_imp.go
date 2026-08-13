package usecase

import (
	"errors"
	"strings"

	"p2-ip-kartikajls/internal/dto"
	"p2-ip-kartikajls/internal/entity"
	"p2-ip-kartikajls/internal/repository"
)

type cameraUsecase struct {
	cameraRepository repository.CameraRepository
}

func NewCameraUsecase(
	cameraRepository repository.CameraRepository,
) CameraUsecase {
	return &cameraUsecase{
		cameraRepository: cameraRepository,
	}
}

// ADMIN - CREATE CAMERA
func (u *cameraUsecase) Create(request dto.CreateCameraRequest) (*dto.CameraResponse, error) {

	name := strings.TrimSpace(request.Name)
	category := strings.TrimSpace(request.Category)

	if name == "" {
		return nil, errors.New("camera name is required")
	}

	if category == "" {
		return nil, errors.New("camera category is required")
	}

	if request.RentalCost <= 0 {
		return nil, errors.New(
			"rental cost must be greater than 0",
		)
	}

	camera := &entity.Camera{
		Name:       name,
		Available:  true,
		RentalCost: request.RentalCost,
		Category:   category,
	}

	err := u.cameraRepository.Create(camera)

	if err != nil {
		return nil, err
	}

	return cameraToResponse(camera), nil
}

// USER + ADMIN - GET ALL
func (u *cameraUsecase) GetAll() ([]dto.CameraResponse, error) {

	cameras, err := u.cameraRepository.GetAll()

	if err != nil {
		return nil, err
	}

	responses := make(
		[]dto.CameraResponse,
		0,
		len(cameras),
	)

	for _, camera := range cameras {

		responses = append(
			responses,
			*cameraToResponse(&camera),
		)
	}

	return responses, nil
}

// USER + ADMIN - GET BY ID
func (u *cameraUsecase) GetByID(cameraID int64) (*dto.CameraResponse, error) {

	if cameraID <= 0 {
		return nil, errors.New("invalid camera id")
	}

	camera, err := u.cameraRepository.GetByID(cameraID)

	if err != nil {
		return nil, errors.New("camera not found")
	}

	return cameraToResponse(camera), nil
}

// ADMIN - UPDATE
func (u *cameraUsecase) Update(cameraID int64, request dto.UpdateCameraRequest) (*dto.CameraResponse, error) {

	if cameraID <= 0 {
		return nil, errors.New("invalid camera id")
	}

	name := strings.TrimSpace(request.Name)
	category := strings.TrimSpace(request.Category)

	if name == "" {
		return nil, errors.New("camera name is required")
	}

	if category == "" {
		return nil, errors.New("camera category is required")
	}

	if request.RentalCost <= 0 {
		return nil, errors.New(
			"rental cost must be greater than 0",
		)
	}

	camera, err := u.cameraRepository.GetByID(cameraID)

	if err != nil {
		return nil, errors.New("camera not found")
	}

	camera.Name = name
	camera.RentalCost = request.RentalCost
	camera.Category = category

	if request.Available != nil {
		camera.Available = *request.Available
	}

	err = u.cameraRepository.Update(camera)

	if err != nil {
		return nil, err
	}

	return cameraToResponse(camera), nil
}

// ADMIN - DELETE
func (u *cameraUsecase) Delete(cameraID int64) error {

	if cameraID <= 0 {
		return errors.New("invalid camera id")
	}

	_, err := u.cameraRepository.GetByID(cameraID)

	if err != nil {
		return errors.New("camera not found")
	}

	return u.cameraRepository.Delete(cameraID)
}

// MAPPER
func cameraToResponse(camera *entity.Camera) *dto.CameraResponse {

	return &dto.CameraResponse{
		CameraID:   camera.CameraID,
		Name:       camera.Name,
		Available:  camera.Available,
		RentalCost: camera.RentalCost,
		Category:   camera.Category,
		CreatedAt:  camera.CreatedAt,
		UpdatedAt:  camera.UpdatedAt,
	}
}
