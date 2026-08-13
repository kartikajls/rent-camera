package repository

import (
	"p2-ip-kartikajls/internal/entity"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(user *entity.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetByID(userID int64) (*entity.User, error) {

	var user entity.User

	err := r.db.
		Where("user_id = ?", userID).
		First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetByEmail(email string) (*entity.User, error) {

	var user entity.User

	err := r.db.
		Where("email = ?", email).
		First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetAll() ([]entity.User, error) {

	var users []entity.User

	err := r.db.
		Order("user_id ASC").
		Find(&users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) Update(user *entity.User) error {
	return r.db.Save(user).Error
}
