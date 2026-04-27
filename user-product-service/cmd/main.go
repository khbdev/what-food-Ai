package main

import (
	"log"
	"user-product-service/internal/cache"
	"user-product-service/internal/config"
	repository "user-product-service/internal/repostory"
	"user-product-service/internal/usecase"
	"user-product-service/pkg/loadenv"
)



func main(){

    loadenv.LoadEnv()

	postgress, err := config.NewPostgresDB()
	if err != nil {
		log.Fatal(err)
	}

	_ = postgress

	redis, err := config.NewRedisClient()
		if err != nil {
		log.Fatal(err)
	}

	_ = redis

	repoCategory := repository.NewCategoryRepository(postgress)
	repoIncrideat := repository.NewCategoryRepository(postgress)

	_ = repoCategory
	_ = repoIncrideat

	cacheCategory := cache.NewCategoryCache(redis)

	_ = cacheCategory

	srvCategory := usecase.NewCategoryUsecase(repoCategory, cacheCategory)

	_ type SortBy []Type
	
	func (a SortBy) Len() int           { return len(a) }
	func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
	func (a SortBy) Less(i, j int) bool { return a[i] < a[j] }
}