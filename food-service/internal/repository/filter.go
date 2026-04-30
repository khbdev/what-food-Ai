package repository

import (
	"context"
	"database/sql"
	"food-service/internal/domain"
	"food-service/internal/models"
	"strings"
)

type foodFilterRepository struct {
	db *sql.DB
}

func NewFoodFilterRepository(db *sql.DB) domain.FoodFilterRepository {
	return &foodFilterRepository{db: db}
}

func (r *foodFilterRepository) Filter(ctx context.Context, filter models.RecipeFilter) ([]models.FoodItemResponse, error) {
	where, args := buildWhere(filter)

	query := `SELECT id, 'recipe' as type, restaurant_id, name, description, image_url, video_url, country, meal_time, kcal, protein, fat, carbs, created_at FROM recipes` + where

	if filter.IncludeSalads {
		saladWhere, saladArgs := buildWhere(filter)
		query += ` UNION ALL SELECT id, 'salad' as type, restaurant_id, name, description, image_url, video_url, country, meal_time, kcal, protein, fat, carbs, created_at FROM salads` + saladWhere
		args = append(args, saladArgs...)
	}

	query += ` ORDER BY created_at DESC`

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.FoodItemResponse
	for rows.Next() {
		var item models.FoodItemResponse
		if err := rows.Scan(
			&item.ID, &item.Type, &item.RestaurantID, &item.Name, &item.Description,
			&item.ImageURL, &item.VideoURL, &item.Country, &item.MealTime,
			&item.Kcal, &item.Protein, &item.Fat, &item.Carbs, &item.CreatedAt,
		); err != nil {
			return nil, err
		}
		results = append(results, item)
	}

	return results, rows.Err()
}

func buildWhere(filter models.RecipeFilter) (string, []any) {
	var conditions []string
	var args []any

	if filter.Country != "" {
		conditions = append(conditions, "country = ?")
		args = append(args, filter.Country)
	}
	if filter.MealTime != "" {
		conditions = append(conditions, "meal_time = ?")
		args = append(args, filter.MealTime)
	}
	if filter.MaxKcal > 0 {
		conditions = append(conditions, "kcal <= ?")
		args = append(args, filter.MaxKcal)
	}

	if len(conditions) == 0 {
		return "", args
	}

	return " WHERE " + strings.Join(conditions, " AND "), args
}