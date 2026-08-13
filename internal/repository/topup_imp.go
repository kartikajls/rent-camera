package repository

import (
	"p2-ip-kartikajls/internal/entity"

	"gorm.io/gorm"
)

type topUpRepository struct {
	db *gorm.DB
}

func NewTopUpRepository(db *gorm.DB) TopUpRepository {
	return &topUpRepository{
		db: db,
	}
}

func (r *topUpRepository) Create(topUp *entity.TopUp) error {
	return r.db.Create(topUp).Error
}

func (r *topUpRepository) GetByID(id int64) (*entity.TopUp, error) {

	var topUp entity.TopUp

	err := r.db.
		Where("top_up_id = ?", id).
		First(&topUp).Error

	if err != nil {
		return nil, err
	}

	return &topUp, nil
}

func (r *topUpRepository) GetByUserID(userID int64) ([]entity.TopUp, error) {

	var topUps []entity.TopUp

	err := r.db.
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&topUps).Error

	return topUps, err
}

func (r *topUpRepository) GetAll() ([]entity.TopUp, error) {

	var topUps []entity.TopUp

	err := r.db.
		Order("created_at DESC").
		Find(&topUps).Error

	return topUps, err
}

func (r *topUpRepository) UpdateStatus(id int64, status string) error {

	return r.db.
		Model(&entity.TopUp{}).
		Where("top_up_id = ?", id).
		Update("status", status).
		Error
}
