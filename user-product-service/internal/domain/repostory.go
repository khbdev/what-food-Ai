package domain


type CategoryRepository interface {
	Create(category *models.Category) error
	GetByID(id int64) (*models.Category, error)
	GetAll() ([]models.Category, error)
	Update(category *models.Category) error
	Delete(id int64) error

	// custom
	GetAllWithUserProducts(userID int64) ([]CategoryWithIngredients, error)
}