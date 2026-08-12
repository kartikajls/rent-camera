package repository

import "p2-ip-kartikajls/internal/entity"

type TopUpRepository interface {
	Create(topUp *entity.TopUp) error
	GetByID(id int64) (*entity.TopUp, error)
	GetByUserID(userID int64) ([]entity.TopUp, error)
	GetAll() ([]entity.TopUp, error)
	UpdateStatus(id int64, status string) error
}
