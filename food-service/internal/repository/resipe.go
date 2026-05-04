package repository

import (
	"context"
	"database/sql"
	"food-service/internal/domain"
	"food-service/internal/models"
	"time"
)

type recipeRepository struct {
	db *sql.DB
}

func NewRecipeRepository(db *sql.DB) domain.RecipeRepository {
	return &recipeRepository{db: db}
}

// =======================
// CREATE
// =======================
func (r *recipeRepository) Create(ctx context.Context, recipe *models.Recipe) error {
	query := `
		INSERT INTO recipes (
			restaurant_id, name, description, image_url, video_url,
			country, meal_time, kcal, protein, fat, carbs, created_at
		)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)
	`

	recipe.CreatedAt = time.Now()

	_, err := r.db.ExecContext(ctx, query,
		recipe.RestaurantID,
		recipe.Name,
		recipe.Description,
		recipe.ImageURL,
		recipe.VideoURL,
		recipe.Country,
		recipe.MealTime,
		recipe.Kcal,
		recipe.Protein,
		recipe.Fat,
		recipe.Carbs,
		recipe.CreatedAt,
	)

	return err
}

// =======================
// GET BY ID
// =======================
func (r *recipeRepository) GetByID(ctx context.Context, id int64) (*models.Recipe, error) {
	query := `
		SELECT id, restaurant_id, name, description, image_url, video_url,
		       country, meal_time, kcal, protein, fat, carbs, created_at
		FROM recipes
		WHERE id = $1
	`

	var recipe models.Recipe

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&recipe.ID,
		&recipe.RestaurantID,
		&recipe.Name,
		&recipe.Description,
		&recipe.ImageURL,
		&recipe.VideoURL,
		&recipe.Country,
		&recipe.MealTime,
		&recipe.Kcal,
		&recipe.Protein,
		&recipe.Fat,
		&recipe.Carbs,
		&recipe.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &recipe, nil
}

// =======================
// GET ALL
// =======================
func (r *recipeRepository) GetAll(ctx context.Context) ([]*models.Recipe, error) {
	query := `
		SELECT id, restaurant_id, name, description, image_url, video_url,
		       country, meal_time, kcal, protein, fat, carbs, created_at
		FROM recipes
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var recipes []*models.Recipe

	for rows.Next() {
		var recipe models.Recipe

		if err := rows.Scan(
			&recipe.ID,
			&recipe.RestaurantID,
			&recipe.Name,
			&recipe.Description,
			&recipe.ImageURL,
			&recipe.VideoURL,
			&recipe.Country,
			&recipe.MealTime,
			&recipe.Kcal,
			&recipe.Protein,
			&recipe.Fat,
			&recipe.Carbs,
			&recipe.CreatedAt,
		); err != nil {
			return nil, err
		}

		recipes = append(recipes, &recipe)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return recipes, nil
}

// =======================
// UPDATE
// =======================
func (r *recipeRepository) Update(ctx context.Context, recipe *models.Recipe) error {
	query := `
		UPDATE recipes
		SET name=$1, description=$2, image_url=$3, video_url=$4,
		    country=$5, meal_time=$6, kcal=$7, protein=$8, fat=$9, carbs=$10
		WHERE id=$11
	`

	_, err := r.db.ExecContext(ctx, query,
		recipe.Name,
		recipe.Description,
		recipe.ImageURL,
		recipe.VideoURL,
		recipe.Country,
		recipe.MealTime,
		recipe.Kcal,
		recipe.Protein,
		recipe.Fat,
		recipe.Carbs,
		recipe.ID,
	)

	return err
}

// =======================
// DELETE
// =======================
func (r *recipeRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM recipes WHERE id=$1`

	_, err := r.db.ExecContext(ctx, query, id)
	return err
}