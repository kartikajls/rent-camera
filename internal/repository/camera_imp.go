package repository

import (
	"p2-ip-kartikajls/internal/entity"

	"gorm.io/gorm"
)

type cameraRepository struct {
	db *gorm.DB
}

func NewCameraRepository(db *gorm.DB) CameraRepository {
	return &cameraRepository{
		db: db,
	}
}

// CREATE
func (r *cameraRepository) Create(camera *entity.Camera) error {

	return r.db.Create(camera).Error
}

// GET ALL
func (r *cameraRepository) GetAll() ([]entity.Camera, error) {

	var cameras []entity.Camera

	err := r.db.
		Order("camera_id ASC").
		Find(&cameras).
		Error

	return cameras, err
}

// GET BY ID
func (r *cameraRepository) GetByID(cameraID int64) (*entity.Camera, error) {

	var camera entity.Camera

	err := r.db.
		Where("camera_id = ?", cameraID).
		First(&camera).
		Error

	if err != nil {
		return nil, err
	}

	return &camera, nil
}

// UPDATE
func (r *cameraRepository) Update(camera *entity.Camera) error {

	return r.db.
		Model(&entity.Camera{}).
		Where("camera_id = ?", camera.CameraID).
		Updates(map[string]interface{}{
			"name":        camera.Name,
			"available":   camera.Available,
			"rental_cost": camera.RentalCost,
			"category":    camera.Category,
		}).Error
}

// DELETE
func (r *cameraRepository) Delete(cameraID int64) error {

	return r.db.
		Where("camera_id = ?", cameraID).
		Delete(&entity.Camera{}).
		Error
}
