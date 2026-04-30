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

// constructor (DI shu yerda bo‘ladi)
func NewRecipeRepository(db *sql.DB) domain.RecipeRepository {
	return &recipeRepository{
		db: db,
	}
}

// CREATE
func (r *recipeRepository) Create(ctx context.Context, recipe *models.Recipe) error {
	query := `
		INSERT INTO recipes (
			restaurant_id, name, description, image_url, video_url,
			country, meal_time, is_salad, kcal, protein, fat, carbs, created_at
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	recipe.CreatedAt = time.Now()

	_, err := r.db.ExecContext(
		ctx,
		query,
		recipe.RestaurantID,
		recipe.Name,
		recipe.Description,
		recipe.ImageURL,
		recipe.VideoURL,
		recipe.Country,
		recipe.MealTime,
		recipe.
		recipe.Kcal,
		recipe.Protien,
		recipe.Fat,
		recipe.Carbs,
		recipe.CreatedAt,
	)

	return err
}

// GET BY ID
func (r *recipeRepository) GetByID(ctx context.Context, id int64) (*models.Recipe, error) {
	query := `
		SELECT 
			id, restaurant_id, name, description, image_url, video_url,
			country, meal_time, is_salad, kcal, protein, fat, carbs, created_at
		FROM recipes
		WHERE id = ?
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
		&recipe.IsSalad,
		&recipe.Kcal,
		&recipe.Protien,
		&recipe.Fat,
		&recipe.Carbs,
		&recipe.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &recipe, nil
}

// GET ALL
func (r *recipeRepository) GetAll(ctx context.Context) ([]*models.Recipe, error) {
	query := `
		SELECT 
			id, restaurant_id, name, description, image_url, video_url,
			country, meal_time, is_salad, kcal, protein, fat, carbs, created_at
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

		err := rows.Scan(
			&recipe.ID,
			&recipe.RestaurantID,
			&recipe.Name,
			&recipe.Description,
			&recipe.ImageURL,
			&recipe.VideoURL,
			&recipe.Country,
			&recipe.MealTime,
			&recipe.IsSalad,
			&recipe.Kcal,
			&recipe.Protien,
			&recipe.Fat,
			&recipe.Carbs,
			&recipe.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		recipes = append(recipes, &recipe)
	}

	return recipes, nil
}

// UPDATE
func (r *recipeRepository) Update(ctx context.Context, recipe *models.Recipe) error {
	query := `
		UPDATE recipes
		SET 
			name = ?,
			description = ?,
			image_url = ?,
			video_url = ?,
			country = ?,
			meal_time = ?,
			is_salad = ?,
			kcal = ?,
			protein = ?,
			fat = ?,
			carbs = ?
		WHERE id = ?
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		recipe.Name,
		recipe.Description,
		recipe.ImageURL,
		recipe.VideoURL,
		recipe.Country,
		recipe.MealTime,
		recipe.IsSalad,
		recipe.Kcal,
		recipe.Protien,
		recipe.Fat,
		recipe.Carbs,
		recipe.ID,
	)

	return err
}

// DELETE
func (r *recipeRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM recipes WHERE id = ?`

	_, err := r.db.ExecContext(ctx, query, id)
	return err
}