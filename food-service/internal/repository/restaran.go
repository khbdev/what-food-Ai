package repository

import (
	"context"
	"database/sql"
	"food-service/internal/domain"
	"food-service/internal/models"
)

type restaurantRepository struct {
	db *sql.DB
}

func NewRestaurantRepository(db *sql.DB) domain.RestaurantRepository {
	return &restaurantRepository{db: db}
}

// =======================
// CREATE
// =======================
func (r *restaurantRepository) Create(ctx context.Context, rest *models.Restaurant) (int64, error) {
	query := `
		INSERT INTO restaurants (restaurant_name, description, image_url, created_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	var id int64

	err := r.db.QueryRowContext(ctx, query,
		rest.RestaurantName,
		rest.Description,
		rest.ImageURL,
		rest.CreatedAt,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

// =======================
// GET BY ID
// =======================
func (r *restaurantRepository) GetByID(ctx context.Context, id int64) (*models.Restaurant, error) {
	if id <= 0 {
		return nil, sql.ErrNoRows
	}

	query := `
		SELECT id, restaurant_name, description, image_url, created_at
		FROM restaurants
		WHERE id = $1
	`

	var rest models.Restaurant

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&rest.ID,
		&rest.RestaurantName,
		&rest.Description,
		&rest.ImageURL,
		&rest.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &rest, nil
}

// =======================
// GET ALL
// =======================
func (r *restaurantRepository) GetAll(ctx context.Context) ([]*models.Restaurant, error) {
	query := `
		SELECT id, restaurant_name, description, image_url, created_at
		FROM restaurants
		ORDER BY id DESC
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var restaurants []*models.Restaurant

	for rows.Next() {
		var rest models.Restaurant

		err := rows.Scan(
			&rest.ID,
			&rest.RestaurantName,
			&rest.Description,
			&rest.ImageURL,
			&rest.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		restaurants = append(restaurants, &rest)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return restaurants, nil
}

// =======================
// UPDATE
// =======================
func (r *restaurantRepository) Update(ctx context.Context, rest *models.Restaurant) error {
	if rest.ID <= 0 {
		return sql.ErrNoRows
	}

	query := `
		UPDATE restaurants
		SET restaurant_name = $1, description = $2, image_url = $3
		WHERE id = $4
	`

	_, err := r.db.ExecContext(ctx, query,
		rest.RestaurantName,
		rest.Description,
		rest.ImageURL,
		rest.ID,
	)

	return err
}

// =======================
// DELETE
// =======================
func (r *restaurantRepository) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return sql.ErrNoRows
	}

	query := `DELETE FROM restaurants WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, id)
	return err
}