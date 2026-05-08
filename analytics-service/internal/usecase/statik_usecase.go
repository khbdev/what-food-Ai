package usecase

import (
	"errors"
	"strings"


)

type MealUsecase struct {
	repo *repository.MealRepository
}

// DI
func NewMealUsecase(repo *repository.MealRepository) *MealUsecase {
	return &MealUsecase{
		repo: repo,
	}
}

// Create meal
func (u *MealUsecase) Create(meal *models.Meal) error {

	// =====================
	// VALIDATION
	// =====================

	if meal.UserID == 0 {
		return errors.New("user_id is required")
	}

	if strings.TrimSpace(meal.Name) == "" {
		return errors.New("name is required")
	}

	if len(meal.Name) < 2 {
		return errors.New("name is too short")
	}

	if strings.TrimSpace(meal.Country) == "" {
		return errors.New("country is required")
	}

	if strings.TrimSpace(meal.MealTime) == "" {
		return errors.New("meal_time is required")
	}

	// meal time check
	validMealTimes := map[string]bool{
		"breakfast": true,
		"lunch":     true,
		"dinner":    true,
		"snack":     true,
	}

	if !validMealTimes[strings.ToLower(meal.MealTime)] {
		return errors.New("invalid meal_time")
	}

	// nutrition validation
	if meal.Kcal < 0 {
		return errors.New("kcal cannot be negative")
	}

	if meal.Protein < 0 {
		return errors.New("protein cannot be negative")
	}

	if meal.Fat < 0 {
		return errors.New("fat cannot be negative")
	}

	if meal.Carbs < 0 {
		return errors.New("carbs cannot be negative")
	}

	// =====================
	// SAVE
	// =====================

	return u.repo.Create(meal)
}