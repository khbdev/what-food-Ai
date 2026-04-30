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

func NewRestaurantRepository(db *sql.DB) domain. {
	return &restaurantRepository{db: db}
}