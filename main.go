package main

import (
	"github.com/Cagge/models"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	models.ConnectDB()
	/*	route.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "Hello World!"})
	})*/
	// Маршруты
	//	route.GET("/tracks", controllers.GetAllTracks)
	route.Run()
}
