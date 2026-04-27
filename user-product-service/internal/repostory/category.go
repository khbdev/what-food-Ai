package repostory

import (
	"database/sql"
	"user-product-service/internal/domain"
	"user-product-service/internal/models"
)




type categoryRepo struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) domain.CategoryRepository {
	return &categoryRepo{db: db}
}


func (r *categoryRepo) Create(c *models.Category) error {
	query := `INSERT INTO categories (name, created_at)
			  VALUES ($1, NOW())
			  RETURNING id, created_at`

	return r.db.QueryRow(query, c.Name).
		Scan(&c.ID, &c.CreatedAt)
}


func (r *categoryRepo) GetByID(id int64) (*models.Category, error) {
	query := `SELECT id, name, created_at FROM categories WHERE id = $1`

	var c models.Category

	err := r.db.QueryRow(query, id).
		Scan(&c.ID, &c.Name, &c.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &c, nil
}