package routes

import (
	controller "github.com/Mekde-lawit/Restaurant-Management/src/controllers"
	"github.com/gin-gonic/gin"
)

func FoodRoutes(route *gin.Engine) {
	route.GET("/foods", controller.GetFoods)
	route.GET("/food/:id", controller.GetFood)
	route.POST("/food", controller.CreateFood)
	route.PATCH("/food/:id", controller.UpdateFood)
}
