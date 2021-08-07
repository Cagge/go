package main

import (
	"github.com/Cagge/controllers"
	"github.com/Cagge/models"
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	models.ConnectDB()
	route.GET("/accs", controllers.GetAllAcc)
	route.POST("/accs", controllers.CreateAcc)
	route.GET("/accs/:id", controllers.GetAcc)
	route.PATCH("/accs/:id", controllers.UpdateAcc)
	route.DELETE("/accs/:id", controllers.DeleteAcc)
	route.Run()
}
