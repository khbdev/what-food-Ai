package repository

import (
	"analytics-service/internal/models"
	"context"
	"time"

	"gorm.io/gorm"
)

type MealRepository struct {
	db *gorm.DB
}

type DailyNutrition struct {
	Day     string  `json:"day"`
	Kcal    float64 `json:"kcal"`
	Fat     float64 `json:"fat"`
	Carbs   float64 `json:"carbs"`
	Protein float64 `json:"protein"`
}

func NewMealRepository(db *gorm.DB) *MealRepository {
	return &MealRepository{
		db: db,
	}
}

func (r *MealRepository) Create(meal *models.Meal) error {
	return r.db.Create(meal).Error
}

func (r *MealRepository) GetWeeklyNutrition(ctx context.Context, userID uint) ([]DailyNutrition, error) {

	now := time.Now()
	weekday := int(now.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	weekStart := now.AddDate(0, 0, -(weekday - 1)).Truncate(24 * time.Hour)
	weekEnd := weekStart.AddDate(0, 0, 7)

	var results []DailyNutrition

	err := r.db.WithContext(ctx).
		Model(&models.Meal{}).
		Select(`
			TO_CHAR(meal_time, 'Day') as day,
			SUM(kcal)    as kcal,
			SUM(fat)     as fat,
			SUM(carbs)   as carbs,
			SUM(protein) as protein
		`).
		Where("user_id = ? AND meal_time >= ? AND meal_time < ?", userID, weekStart, weekEnd).
		Group("TO_CHAR(meal_time, 'Day'), DATE(meal_time)").
		Order("DATE(meal_time)").
		Scan(&results).Error

	return results, err
}