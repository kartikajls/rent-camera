package entity

import "time"

type EmailNotification struct {
	NotificationID    int64      `gorm:"primaryKey;column:notification_id"`
	UserID            int64      `gorm:"not null;column:user_id"`
	RentalOrderID     *int64     `gorm:"column:rental_order_id"`
	Email             string     `gorm:"type:varchar(255);not null"`
	NotificationType  string     `gorm:"type:varchar(50);not null;column:notification_type"`
	Subject           string     `gorm:"type:varchar(255);not null"`
	Status            string     `gorm:"type:varchar(30);not null;default:'pending'"`
	Provider          *string    `gorm:"type:varchar(50);column:provider"`
	ProviderMessageID *string    `gorm:"type:varchar(255);column:provider_message_id"`
	ErrorMessage      *string    `gorm:"type:text;column:error_message"`
	SentAt            *time.Time `gorm:"column:sent_at"`
	CreatedAt         time.Time  `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt         time.Time  `gorm:"column:updated_at;autoUpdateTime"`
}

func (EmailNotification) TableName() string {
	return "email_notifications"
}
