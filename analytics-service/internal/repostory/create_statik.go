package repository

import (


	"gorm.io/gorm"
)

type MealRepository struct {
	db *gorm.DB
}

// DI pattern
func NewMealRepository(db *gorm.DB) *MealRepository {
	return &MealRepository{
		db: db,
	}
}

// Create meal
func (r *MealRepository) Create(meal *mod.Meal) error {
	return r.db.Create(meal).Error
}