package repository

import (
	"context"
	"database/sql"
	"fmt"
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

func (r *foodFilterRepository) Filter(
	ctx context.Context,
	filter models.RecipeFilter,
) ([]models.FoodItemResponse, error) {

	recipeWhere, recipeArgs := buildWhere(filter)

	query := `
		SELECT id, 'recipe' as type, restaurant_id, name, description,
		       image_url, video_url, country, meal_time,
		       kcal, protein, fat, carbs, created_at
		FROM recipes
	` + recipeWhere

	args := recipeArgs

	if filter.IncludeSalads {
		saladWhere, saladArgs := buildWhere(filter)

		query += `
			UNION ALL
			SELECT id, 'salad' as type, restaurant_id, name, description,
			       image_url, video_url, country, meal_time,
			       kcal, protein, fat, carbs, created_at
			FROM salads
		` + saladWhere

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

		err := rows.Scan(
			&item.ID,
			&item.Type,
			&item.RestaurantID,
			&item.Name,
			&item.Description,
			&item.ImageURL,
			&item.VideoURL,
			&item.Country,
			&item.MealTime,
			&item.Kcal,
			&item.Protein,
			&item.Fat,
			&item.Carbs,
			&item.CreatedAt,
		)
		if err != nil {
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
		args = append(args, filter.Country)
		conditions = append(conditions, fmt.Sprintf("country = $%d", len(args)))
	}

	if filter.MealTime != "" {
		args = append(args, filter.MealTime)
		conditions = append(conditions, fmt.Sprintf("meal_time = $%d", len(args)))
	}

	if filter.MaxKcal > 0 {
		args = append(args, filter.MaxKcal)
		conditions = append(conditions, fmt.Sprintf("kcal <= $%d", len(args)))
	}

	if len(conditions) == 0 {
		return "", nil
	}

	return " WHERE " + strings.Join(conditions, " AND "), args
}