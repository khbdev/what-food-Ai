package repository

import (
	"context"
	"user-service/internal/domain"
	"user-service/internal/models"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// DI constructor
func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{
		db: db,
	}
}

// CREATE
func (r *userRepository) Create(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

// GET BY ID
func (r *userRepository) GetByID(ctx context.Context, id uint) (*models.User, error) {
	var user models.User

	err := r.db.WithContext(ctx).
		First(&user, id).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// GET BY PHONE
func (r *userRepository) GetByPhone(ctx context.Context, phone string) (*models.User, error) {
	var user models.User

	err := r.db.WithContext(ctx).
		Where("phone = ?", phone).
		First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// GET ALL
func (r *userRepository) GetAll(ctx context.Context, limit, offset int) ([]models.User, error) {
	var users []models.User

	err := r.db.WithContext(ctx).
		Limit(limit).
		Offset(offset).
		Find(&users).Error

	return users, err
}

// UPDATE
func (r *userRepository) Update(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).
		Save(user).Error
}

// DELETE
func (r *userRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).
		Delete(&models.User{}, id).Error
}