package repository

import "p2-ip-kartikajls/internal/entity"

type UserRepository interface {
	Create(user *entity.User) error
	GetByID(userID int64) (*entity.User, error)
	GetByEmail(email string) (*entity.User, error)
	GetAll() ([]entity.User, error)
	Update(user *entity.User) error
}
