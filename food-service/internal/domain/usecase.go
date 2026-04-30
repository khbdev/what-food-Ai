package domain



type RecipeUsecase interface {
	Create(ctx context.Context, recipe *models.Recipe) error
	GetByID(ctx context.Context, id int64) (*models.Recipe, error)
	GetAll(ctx context.Context) ([]*models.Recipe, error)
	Update(ctx context.Context, recipe *models.Recipe) error
	Delete(ctx context.Context, id int64) error
}