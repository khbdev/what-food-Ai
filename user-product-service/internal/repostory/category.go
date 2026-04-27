package repostory

import (
	"database/sql"
	"user-product-service/internal/domain"
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