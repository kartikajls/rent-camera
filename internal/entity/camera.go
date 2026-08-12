package entity

import "time"

type Camera struct {
	CameraID   int64     `gorm:"primaryKey;column:camera_id"`
	Name       string    `gorm:"type:varchar(255);not null"`
	Available  bool      `gorm:"not null;default:true"`
	RentalCost float64   `gorm:"type:numeric(15,2);not null"`
	Category   string    `gorm:"type:varchar(100);not null"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt  time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (Camera) TableName() string {
	return "cameras"
}
