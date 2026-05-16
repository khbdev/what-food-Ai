package client

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	categorypb "github.com/khbdev/what-food-proto/proto/products"
	ingredientpb "github.com/khbdev/what-food-proto/proto/incrideats"
)

// =========================
// STRUCT
// =========================

type UserProductClient struct {
	conn *grpc.ClientConn

	Category   categorypb.CategoryServiceClient
	Ingredient ingredientpb.IngredientServiceClient
}

// =========================
// CONFIG
// =========================

const timeoutUSERPRODUCT = 5 * time.Second

// =========================
// INIT (FAIL FAST)
// =========================

func NewUserProductClient(addr string) (*UserProductClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutUSERPRODUCT)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, err
	}

	log.Println("✅ User Product service connected:", addr)

	return &UserProductClient{
		conn:       conn,
		Category:   categorypb.NewCategoryServiceClient(conn),
		Ingredient: ingredientpb.NewIngredientServiceClient(conn),
	}, nil
}

// =========================
// CLOSE
// =========================

func (c *UserProductClient) Close() error {
	return c.conn.Close()
}

// =========================
// CONTEXT HELPER
// =========================

func (c *UserProductClient) ctx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), timeoutUSERPRODUCT)
}

//
// =====================================
// CATEGORY METHODS
// =====================================
//

func (c *UserProductClient) CreateCategory(req *categorypb.CreateCategoryRequest) (*categorypb.CategoryResponse, error) {
	ctx, cancel := c.ctx()
	defer cancel()

	return c.Category.Create(ctx, req)
}

func (c *UserProductClient) GetCategoryByID(req *categorypb.GetByIDRequest) (*categorypb.CategoryWithIngredientsResponse, error) {
	ctx, cancel := c.ctx()
	defer cancel()

	return c.Category.GetByID(ctx, req)
}

func (c *UserProductClient) GetAllCategories(req *categorypb.GetAllRequest) (*categorypb.CategoryListResponse, error) {
	ctx, cancel := c.ctx()
	defer cancel()

	return c.Category.GetAll(ctx, req)
}

func (c *UserProductClient) UpdateCategory(req *categorypb.UpdateCategoryRequest) (*categorypb.CategoryResponse, error) {
	ctx, cancel := c.ctx()
	defer cancel()

	return c.Category.Update(ctx, req)
}

func (c *UserProductClient) DeleteCategory(req *categorypb.DeleteCategoryRequest) (*categorypb.Empty, error) {
	ctx, cancel := c.ctx()
	defer cancel()

	return c.Category.Delete(ctx, req)
}

func (c *UserProductClient) GetAllWithUserProducts(req *categorypb.GetAllWithUserProductsRequest) (*categorypb.CategoryWithIngredientsListResponse, error) {
	ctx, cancel := c.ctx()
	defer cancel()

	return c.Category.GetAllWithUserProducts(ctx, req)
}

func (c *UserProductClient) GetCategoryByIDWithUserProducts(req *categorypb.GetByIDWithUserProductsRequest) (*categorypb.CategoryWithIngredientsResponse, error) {
	ctx, cancel := c.ctx()
	defer cancel()

	return c.Category.GetByIDWithUserProducts(ctx, req)
}

//
// =====================================
// INGREDIENT METHODS
// =====================================
//

func (c *UserProductClient) CreateIngredient(req *ingredientpb.CreateIngredientRequest) (*ingredientpb.IngredientResponse, error) {
	ctx, cancel := c.ctx()
	defer cancel()

	return c.Ingredient.Create(ctx, req)
}

func (c *UserProductClient) GetIngredientByID(req *ingredientpb.GetIngredientByIDRequest) (*ingredientpb.IngredientResponse, error) {
	ctx, cancel := c.ctx()
	defer cancel()

	return c.Ingredient.GetByID(ctx, req)
}

func (c *UserProductClient) GetAllIngredients(req *ingredientpb.GetAllIngredientsRequest) (*ingredientpb.IngredientListResponse, error) {
	ctx, cancel := c.ctx()
	defer cancel()

	return c.Ingredient.GetAll(ctx, req)
}

func (c *UserProductClient) UpdateIngredient(req *ingredientpb.UpdateIngredientRequest) (*ingredientpb.IngredientResponse, error) {
	ctx, cancel := c.ctx()
	defer cancel()

	return c.Ingredient.Update(ctx, req)
}

func (c *UserProductClient) DeleteIngredient(req *ingredientpb.DeleteIngredientRequest) (*ingredientpb.Empty, error) {
	ctx, cancel := c.ctx()
	defer cancel()

	return c.Ingredient.Delete(ctx, req)
}