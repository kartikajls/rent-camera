package entity

import "time"

type RentalOrderDetail struct {
	RentalDetailID int64     `gorm:"primaryKey;column:rental_detail_id"`
	RentalOrderID  int64     `gorm:"not null;column:rental_order_id"`
	CameraID       int64     `gorm:"not null;column:camera_id"`
	RentalCost     float64   `gorm:"type:numeric(15,2);not null;column:rental_cost"`
	RentalDays     int       `gorm:"not null;default:1;column:rental_days"`
	Subtotal       float64   `gorm:"type:numeric(15,2);not null;column:subtotal"`
	CreatedAt      time.Time `gorm:"not null;column:created_at"`
}

func (RentalOrderDetail) TableName() string {
	return "rental_order_details"
}
