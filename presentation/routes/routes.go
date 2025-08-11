package routes

import (
	docs "github.com/dornascarol/api-go-gin/docs"
	"github.com/dornascarol/api-go-gin/presentation/controllers"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func HandleRequests(
    singerController *controllers.SingerController,
    createSingerController *controllers.CreateSingerController,
	searchSingerByIdController *controllers.SearchSingerByIdController,
	deleteSingerController *controllers.DeleteSingerController,
	editSingerController *controllers.EditSingerController,
	searchSingerByNameController *controllers.SearchSingerByNameController,
	greetingController *controllers.GreetingController,
	indexPageController *controllers.IndexPageController,
	notFoundController *controllers.NotFoundController,
) *gin.Engine {
	
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/singers/v1"
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")

	r.GET("/singers", singerController.GetSingers)
	r.POST("/singers", createSingerController.CreateNewSinger)
	r.GET("/singers/:id", searchSingerByIdController.SearchSingerById)
	r.DELETE("/singers/:id", deleteSingerController.DeleteSinger)
	r.PATCH("/singers/:id", editSingerController.EditSinger)
	r.GET("/singers/name/:name", searchSingerByNameController.SearchSingerByName)
	r.GET("/:name", greetingController.Greeting)
	r.GET("/index", indexPageController.DisplayIndexPage)
	r.NoRoute(notFoundController.RouteNotFound)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}
