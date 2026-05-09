package client

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	foodpb "github.com/khbdev/what-food-proto/proto/food"
)

const timeout = 5 * time.Second

type FoodClient struct {
	conn *grpc.ClientConn

	Recipe     foodpb.RecipeServiceClient
	Salad      foodpb.SaladServiceClient
	Restaurant foodpb.RestaurantServiceClient
	Filter     foodpb.FoodFilterServiceClient
}

func NewFoodClient(addr string) (*FoodClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
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

	log.Println("✅ Food service connected:", addr)

	return &FoodClient{
		conn:       conn,
		Recipe:     foodpb.NewRecipeServiceClient(conn),
		Salad:      foodpb.NewSaladServiceClient(conn),
		Restaurant: foodpb.NewRestaurantServiceClient(conn),
		Filter:     foodpb.NewFoodFilterServiceClient(conn),
	}, nil
}

func (c *FoodClient) Close() error {
	return c.conn.Close()
}

func (c *FoodClient) ctx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), timeout)
}

// ================== RECIPE ==================

func (c *FoodClient) CreateRecipe(req *foodpb.CreateRecipeRequest) error {
	ctx, cancel := c.ctx()
	defer cancel()

	_, err := c.Recipe.CreateRecipe(ctx, req)
	return err
}

func (c *FoodClient) GetRecipeByID(id int64) (*foodpb.Recipe, error) {
	ctx, cancel := c.ctx()
	defer cancel()

	res, err := c.Recipe.GetRecipeByID(ctx, &foodpb.GetByIDRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return res.Recipe, nil
}

func (c *FoodClient) GetAllRecipes() ([]*foodpb.Recipe, error) {
	ctx, cancel := c.ctx()
	defer cancel()

	res, err := c.Recipe.GetAllRecipes(ctx, &foodpb.Empty{})
	if err != nil {
		return nil, err
	}
	return res.Recipes, nil
}

func (c *FoodClient) UpdateRecipe(req *foodpb.UpdateRecipeRequest) error {
	ctx, cancel := c.ctx()
	defer cancel()

	_, err := c.Recipe.UpdateRecipe(ctx, req)
	return err
}

func (c *FoodClient) DeleteRecipe(id int64) error {
	ctx, cancel := c.ctx()
	defer cancel()

	_, err := c.Recipe.DeleteRecipe(ctx, &foodpb.GetByIDRequest{Id: id})
	return err
}

// ================== SALAD ==================

func (c *FoodClient) CreateSalad(req *foodpb.CreateSaladRequest) error {
	ctx, cancel := c.ctx()
	defer cancel()

	_, err := c.Salad.CreateSalad(ctx, req)
	return err
}

func (c *FoodClient) GetSaladByID(id int64) (*foodpb.Salad, error) {
	ctx, cancel := c.ctx()
	defer cancel()

	res, err := c.Salad.GetSaladByID(ctx, &foodpb.GetByIDRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return res.Salad, nil
}

func (c *FoodClient) GetAllSalads() ([]*foodpb.Salad, error) {
	ctx, cancel := c.ctx()
	defer cancel()

	res, err := c.Salad.GetAllSalads(ctx, &foodpb.Empty{})
	if err != nil {
		return nil, err
	}
	return res.Salads, nil
}

func (c *FoodClient) UpdateSalad(req *foodpb.UpdateSaladRequest) error {
	ctx, cancel := c.ctx()
	defer cancel()

	_, err := c.Salad.UpdateSalad(ctx, req)
	return err
}

func (c *FoodClient) DeleteSalad(id int64) error {
	ctx, cancel := c.ctx()
	defer cancel()

	_, err := c.Salad.DeleteSalad(ctx, &foodpb.GetByIDRequest{Id: id})
	return err
}

// ================== RESTAURANT ==================

func (c *FoodClient) CreateRestaurant(req *foodpb.CreateRestaurantRequest) (int64, error) {
	ctx, cancel := c.ctx()
	defer cancel()

	res, err := c.Restaurant.CreateRestaurant(ctx, req)
	if err != nil {
		return 0, err
	}
	return res.Id, nil
}

func (c *FoodClient) GetRestaurantByID(id int64) (*foodpb.Restaurant, error) {
	ctx, cancel := c.ctx()
	defer cancel()

	res, err := c.Restaurant.GetRestaurantByID(ctx, &foodpb.GetByIDRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return res.Restaurant, nil
}

func (c *FoodClient) GetAllRestaurants() ([]*foodpb.Restaurant, error) {
	ctx, cancel := c.ctx()
	defer cancel()

	res, err := c.Restaurant.GetAllRestaurants(ctx, &foodpb.Empty{})
	if err != nil {
		return nil, err
	}
	return res.Restaurants, nil
}

func (c *FoodClient) UpdateRestaurant(req *foodpb.UpdateRestaurantRequest) error {
	ctx, cancel := c.ctx()
	defer cancel()

	_, err := c.Restaurant.UpdateRestaurant(ctx, req)
	return err
}

func (c *FoodClient) DeleteRestaurant(id int64) error {
	ctx, cancel := c.ctx()
	defer cancel()

	_, err := c.Restaurant.DeleteRestaurant(ctx, &foodpb.GetByIDRequest{Id: id})
	return err
}

// ================== FILTER ==================

func (c *FoodClient) FilterFood(req *foodpb.FoodFilterRequest) ([]*foodpb.FoodItem, error) {
	ctx, cancel := c.ctx()
	defer cancel()

	res, err := c.Filter.FilterFood(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.Items, nil
}