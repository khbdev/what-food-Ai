package repository

import (
	"context"
	"database/sql"
	"food-service/internal/models"
)

type restaurantRepository struct {
	db *sql.DB
}

func NewRestaurantRepository(db *sql.DB) RestaurantRepository {
	return &restaurantRepository{db: db}
}