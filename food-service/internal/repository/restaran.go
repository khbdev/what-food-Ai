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

func (r *restaurantRepository) Create(ctx context.Context, rest *models.Restaurant) (int64, error) {
	query := `
		INSERT INTO restaurants (restaurant_name, description, image_url, created_at)
		VALUES (?, ?, ?, ?)
	`

	res, err := r.db.ExecContext(ctx, query,
		rest.RestaurantName,
		rest.Description,
		rest.ImageURL,
		rest.CreatedAt,
	)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}