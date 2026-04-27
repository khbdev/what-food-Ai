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

func (r *categoryRepo) GetAll() ([]models.Category, error) {
	query := `SELECT id, name, created_at FROM categories`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []models.Category

	for rows.Next() {
		var c models.Category
		if err := rows.Scan(&c.ID, &c.Name, &c.CreatedAt); err != nil {
			return nil, err
		}
		res = append(res, c)
	}

	return res, nil
}


func (r *categoryRepo) Update(c *models.Category) error {
	query := `UPDATE categories SET name=$1 WHERE id=$2`

	_, err := r.db.Exec(query, c.Name, c.ID)
	return err
}