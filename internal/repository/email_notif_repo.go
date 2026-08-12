package repository

import (
	"p2-ip-kartikajls/internal/entity"
)

type EmailNotificationRepository interface {
	Create(notification *entity.EmailNotification) error
	GetByID(notificationID int64) (*entity.EmailNotification, error)
	GetByUserID(userID int64) ([]entity.EmailNotification, error)
	Update(notification *entity.EmailNotification) error
}
