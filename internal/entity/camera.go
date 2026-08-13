package entity

import "time"

type Camera struct {
	CameraID   int64     `gorm:"primaryKey;column:camera_id" json:"camera_id"`
	Name       string    `gorm:"type:varchar(150);not null;column:name" json:"name"`
	Available  bool      `gorm:"not null;default:true;column:available" json:"available"`
	RentalCost float64   `gorm:"type:numeric(15,2);not null;column:rental_cost" json:"rental_cost"`
	Category   string    `gorm:"type:varchar(100);not null;column:category" json:"category"`
	CreatedAt  time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;column:created_at" json:"created_at"`
	UpdatedAt  time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;column:updated_at" json:"updated_at"`
}

func (Camera) TableName() string {
	return "cameras"
}
