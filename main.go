package main

import (
	"github.com/dornascarol/api-go-gin/application/usecases"
	"github.com/dornascarol/api-go-gin/infrastructure/cache"
	"github.com/dornascarol/api-go-gin/infrastructure/persistence/postgres"
	"github.com/dornascarol/api-go-gin/presentation/controllers"
	"github.com/dornascarol/api-go-gin/presentation/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	db := postgres.ConnectToDatabase()

	redisClient := cache.NewRedisClient()
	cacheService := cache.NewCacheService(redisClient)

	singerRepository := postgres.NewPostgresSingerRepository(db)

	singerUseCase := usecases.NewSingersUseCase(singerRepository, cacheService)
	singerController := controllers.NewSingerController(singerUseCase)
	createSingerController := controllers.NewCreateSingerController(singerUseCase)
	searchSingerByIdController := controllers.NewSearchSingerByIdController(singerUseCase)
	deleteSingerController := controllers.NewDeleteSingerController(singerUseCase)
	editSingerController := controllers.NewEditSingerController(singerUseCase)
	searchSingerByNameController := controllers.NewSearchSingerByNameController(singerUseCase)
	greetingController := controllers.NewGreetingController()
	indexPageController := controllers.NewIndexPageController()
	notFoundController := controllers.NewNotFoundController()

	router := routes.HandleRequests(
		singerController,
		createSingerController,
		searchSingerByIdController,
		deleteSingerController,
		editSingerController,
		searchSingerByNameController,
		greetingController,
		indexPageController,
		notFoundController,
	)
	router.Run()
}
