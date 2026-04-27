package repository

import (
	"database/sql"
	"errors"

	"user-product-service/internal/domain"
	"user-product-service/internal/models"
)

type ingredientRepo struct {
	db *sql.DB
}

func NewIngredientRepository(db *sql.DB) domain.IngredientRepository {
	return &ingredientRepo{db: db}
}

func (r *ingredientRepo) Create(ing *models.Ingredient) error {
	query := `
		INSERT INTO ingredients (user_id, name, quantity, category_id, created_at)
		VALUES ($1, $2, $3, $4, NOW())
		RETURNING id, created_at
	`

	return r.db.QueryRow(
		query,
		ing.UserID,
		ing.Name,
		ing.Quantity,
		ing.CategoryID,
	).Scan(&ing.ID, &ing.CreatedAt)
}

func (r *ingredientRepo) GetByID(id int64, userID int64) (*models.Ingredient, error) {
	query := `
		SELECT id, user_id, name, quantity, category_id, created_at
		FROM ingredients
		WHERE id = $1 AND user_id = $2
	`

	var ing models.Ingredient

	err := r.db.QueryRow(query, id, userID).
		Scan(&ing.ID, &ing.UserID, &ing.Name, &ing.Quantity, &ing.CategoryID, &ing.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("ingredient not found")
		}
		return nil, err
	}

	return &ing, nil
}

func (r *ingredientRepo) GetAll(userID int64) ([]models.Ingredient, error) {
	query := `
		SELECT id, user_id, name, quantity, category_id, created_at
		FROM ingredients
		WHERE user_id = $1
		ORDER BY id DESC
	`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []models.Ingredient

	for rows.Next() {
		var ing models.Ingredient

		if err := rows.Scan(
			&ing.ID,
			&ing.UserID,
			&ing.Name,
			&ing.Quantity,
			&ing.CategoryID,
			&ing.CreatedAt,
		); err != nil {
			return nil, err
		}

		res = append(res, ing)
	}

	return res, nil
}

func (r *ingredientRepo) Update(ing *models.Ingredient) error {
	query := `
		UPDATE ingredients
		SET name=$1, quantity=$2, category_id=$3
		WHERE id=$4 AND user_id=$5
	`

	_, err := r.db.Exec(
		query,
		ing.Name,
		ing.Quantity,
		ing.CategoryID,
		ing.ID,
		ing.UserID,
	)

	return err
}

func (r *ingredientRepo) Delete(id int64, userID int64) error {
	query := `
		DELETE FROM ingredients
		WHERE id=$1 AND user_id=$2
	`

	_, err := r.db.Exec(query, id, userID)
	return err
}