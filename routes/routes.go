package routes

import (
	"github.com/dornascarol/api-go-gin/controllers"
	docs "github.com/dornascarol/api-go-gin/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func HandleRequests() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.GET("/singers", controllers.GetSingers)
	r.POST("/singers", controllers.CreateNewSinger)
	r.GET("/singers/:id", controllers.SearchSingerById)
	r.DELETE("/singers/:id", controllers.DeleteSinger)
	r.PATCH("/singers/:id", controllers.EditSinger)
	r.GET("/singers/name/:name", controllers.SearchSingerByName)
	r.GET("/:name", controllers.Greeting)
	r.GET("/index", controllers.DisplayIndexPage)
	r.NoRoute(controllers.RouteNotFound)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run()
}
