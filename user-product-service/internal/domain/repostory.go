package domain

import "user-product-service/internal/models"


type CategoryRepository interface {
	Create(category *models.Category) error
	GetByID(id int64) (*models.Category, error)
	GetAll() ([]models.Category, error)
	Update(category *models.Category) error
	Delete(id int64) error

	GetAllWithUserProducts(userID int64) ([]mod, error)
}