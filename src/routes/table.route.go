package routes

import (
	controller "github.com/Mekde-lawit/Restaurant-Management/src/controllers"
	"github.com/gin-gonic/gin"
)

func TableRoutes(route *gin.Engine) {
	route.GET("/tables", controller.GetTables)
	route.GET("/table/:id", controller.GetTable)
	route.POST("/table", controller.CreateTable)
	route.PATCH("/table/:id", controller.UpdateTable)
}
