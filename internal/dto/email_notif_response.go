package dto

import "time"

type EmailNotificationResponse struct {
	NotificationID    int64      `json:"notification_id"`
	UserID            int64      `json:"user_id"`
	RentalOrderID     *int64     `json:"rental_order_id,omitempty"`
	Email             string     `json:"email"`
	NotificationType  string     `json:"notification_type"`
	Subject           string     `json:"subject"`
	Status            string     `json:"status"`
	Provider          *string    `json:"provider,omitempty"`
	ProviderMessageID *string    `json:"provider_message_id,omitempty"`
	ErrorMessage      *string    `json:"error_message,omitempty"`
	SentAt            *time.Time `json:"sent_at,omitempty"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
}
