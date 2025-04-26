package routes

import (
	"github.com/dornascarol/api-go-gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/singers", controllers.GetSingers)
	r.GET("/:name", controllers.Greeting)
	r.POST("/singers", controllers.CreateNewSinger)
	r.GET("/singers/:id", controllers.SearchSingerById)
	r.Run()
}
