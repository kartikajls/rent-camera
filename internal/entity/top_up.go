package entity

import "time"

type TopUp struct {
	TopUpID   int64     `gorm:"primaryKey;column:top_up_id"`
	UserID    int64     `gorm:"column:user_id;not null"`
	Amount    float64   `gorm:"type:numeric(15,2);not null"`
	Status    string    `gorm:"type:varchar(20);not null;default:pending"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (TopUp) TableName() string {
	return "top_ups"
}
