package repository

import "p2-ip-kartikajls/internal/entity"

type CameraRepository interface {
	Create(camera *entity.Camera) error
	GetAll() ([]entity.Camera, error)
	GetByID(cameraID int64) (*entity.Camera, error)
	Update(camera *entity.Camera) error
	Delete(cameraID int64) error
}
