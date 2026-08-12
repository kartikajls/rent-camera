package entity

import "time"

type User struct {
	UserID        int64     `gorm:"primaryKey;column:user_id"`
	Username      string    `gorm:"type:varchar(100);not null"`
	Email         string    `gorm:"type:varchar(255);unique;not null"`
	Password      string    `gorm:"type:text;not null"`
	Role          string    `gorm:"type:varchar(20);not null;default:'user'"`
	DepositAmount float64   `gorm:"type:numeric(15,2);not null;default:0"`
	CreatedAt     time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt     time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (User) TableName() string {
	return "users"
}
