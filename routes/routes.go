package routes

import (
	"github.com/dornascarol/api-go-gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
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
	r.Run()
}
