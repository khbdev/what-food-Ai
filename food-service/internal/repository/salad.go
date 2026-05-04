package repository

import (
	"context"
	"database/sql"
	"food-service/internal/domain"
	"food-service/internal/models"
	"time"
)

type saladRepository struct {
	db *sql.DB
}

func NewSaladRepository(db *sql.DB) domain.SaladRepository {
	return &saladRepository{db: db}
}

// =======================
// CREATE
// =======================
func (r *saladRepository) Create(ctx context.Context, salad *models.Salad) error {
	query := `
		INSERT INTO salads (
			restaurant_id, name, description, image_url, video_url,
			country, meal_time, kcal, protein, fat, carbs, created_at
		)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)
	`

	salad.CreatedAt = time.Now()

	_, err := r.db.ExecContext(ctx, query,
		salad.RestaurantID,
		salad.Name,
		salad.Description,
		salad.ImageURL,
		salad.VideoURL,
		salad.Country,
		salad.MealTime,
		salad.Kcal,
		salad.Protein,
		salad.Fat,
		salad.Carbs,
		salad.CreatedAt,
	)

	return err
}

// =======================
// GET BY ID
// =======================
func (r *saladRepository) GetByID(ctx context.Context, id int64) (*models.Salad, error) {
	query := `
		SELECT id, restaurant_id, name, description, image_url, video_url,
		       country, meal_time, kcal, protein, fat, carbs, created_at
		FROM salads
		WHERE id = $1
	`

	var salad models.Salad

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&salad.ID,
		&salad.RestaurantID,
		&salad.Name,
		&salad.Description,
		&salad.ImageURL,
		&salad.VideoURL,
		&salad.Country,
		&salad.MealTime,
		&salad.Kcal,
		&salad.Protein,
		&salad.Fat,
		&salad.Carbs,
		&salad.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &salad, nil
}

// =======================
// GET ALL
// =======================
func (r *saladRepository) GetAll(ctx context.Context) ([]*models.Salad, error) {
	query := `
		SELECT id, restaurant_id, name, description, image_url, video_url,
		       country, meal_time, kcal, protein, fat, carbs, created_at
		FROM salads
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var salads []*models.Salad

	for rows.Next() {
		var s models.Salad

		if err := rows.Scan(
			&s.ID,
			&s.RestaurantID,
			&s.Name,
			&s.Description,
			&s.ImageURL,
			&s.VideoURL,
			&s.Country,
			&s.MealTime,
			&s.Kcal,
			&s.Protein,
			&s.Fat,
			&s.Carbs,
			&s.CreatedAt,
		); err != nil {
			return nil, err
		}

		salads = append(salads, &s)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return salads, nil
}

// =======================
// UPDATE
// =======================
func (r *saladRepository) Update(ctx context.Context, salad *models.Salad) error {
	query := `
		UPDATE salads
		SET name=$1, description=$2, image_url=$3, video_url=$4,
		    country=$5, meal_time=$6, kcal=$7, protein=$8, fat=$9, carbs=$10
		WHERE id=$11
	`

	_, err := r.db.ExecContext(ctx, query,
		salad.Name,
		salad.Description,
		salad.ImageURL,
		salad.VideoURL,
		salad.Country,
		salad.MealTime,
		salad.Kcal,
		salad.Protein,
		salad.Fat,
		salad.Carbs,
		salad.ID,
	)

	return err
}

// =======================
// DELETE
// =======================
func (r *saladRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM salads WHERE id=$1`

	_, err := r.db.ExecContext(ctx, query, id)
	return err
}