package repository

import (
	"p2-ip-kartikajls/internal/entity"

	"gorm.io/gorm"
)

type emailNotificationRepository struct {
	db *gorm.DB
}

func NewEmailNotificationRepository(db *gorm.DB) EmailNotificationRepository {
	return &emailNotificationRepository{
		db: db,
	}
}

// create
func (r *emailNotificationRepository) Create(notification *entity.EmailNotification) error {
	return r.db.Create(notification).Error
}

// Getby Id
func (r *emailNotificationRepository) GetByID(notificationID int64) (*entity.EmailNotification, error) {

	var notification entity.EmailNotification

	err := r.db.
		Where("notification_id = ?", notificationID).
		First(&notification).Error

	if err != nil {
		return nil, err
	}

	return &notification, nil
}

// GetbyUserID
func (r *emailNotificationRepository) GetByUserID(userID int64) ([]entity.EmailNotification, error) {

	var notifications []entity.EmailNotification

	err := r.db.
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&notifications).Error

	if err != nil {
		return nil, err
	}

	return notifications, nil
}

// update
func (r *emailNotificationRepository) Update(notification *entity.EmailNotification) error {
	return r.db.Save(notification).Error
}
