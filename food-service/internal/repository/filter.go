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

// =========================
// MAIN FILTER
// =========================

func (r *foodFilterRepository) Filter(
	ctx context.Context,
	filter models.RecipeFilter,
) ([]models.FoodItemResponse, error) {

	// =========================
	// BASE QUERY (RECIPES)
	// =========================

	recipeWhere, recipeArgs := buildWhere(filter, 1)

	query := `
		SELECT id, 'recipe' as type, restaurant_id, name, description,
		       image_url, video_url, country, meal_time,
		       kcal, protein, fat, carbs, created_at
		FROM recipes
	` + recipeWhere

	args := recipeArgs

	// =========================
	// OPTIONAL SALADS
	// =========================

	if filter.IncludeSalads {

		saladWhere, saladArgs := buildWhere(filter, len(args)+1)

		query += `
			UNION ALL
			SELECT id, 'salad' as type, restaurant_id, name, description,
			       image_url, video_url, country, meal_time,
			       kcal, protein, fat, carbs, created_at
			FROM salads
		` + saladWhere

		args = append(args, saladArgs...)
	}

	// =========================
	// FINAL ORDER
	// =========================

	query += ` ORDER BY created_at DESC`

	// DEBUG (MUHIM)
	// log.Println("SQL:", query)
	// log.Println("ARGS:", args)

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

// =========================
// SAFE POSTGRES WHERE BUILDER
// =========================

func buildWhere(filter models.RecipeFilter, start int) (string, []any) {

	var conditions []string
	var args []any
	i := start

	// country
	if filter.Country != "" {
		i++
		conditions = append(conditions, fmt.Sprintf("country = $%d", i))
		args = append(args, filter.Country)
	}

	// meal_time
	if filter.MealTime != "" {
		i++
		conditions = append(conditions, fmt.Sprintf("meal_time = $%d", i))
		args = append(args, filter.MealTime)
	}

	// kcal
	if filter.MaxKcal > 0 {
		i++
		conditions = append(conditions, fmt.Sprintf("kcal <= $%d", i))
		args = append(args, filter.MaxKcal)
	}

	// ❗ MUHIM FIX: agar hech narsa bo‘lmasa WHERE YO‘Q
	if len(conditions) == 0 {
		return "", []any{}
	}

	return " WHERE " + strings.Join(conditions, " AND "), args
}