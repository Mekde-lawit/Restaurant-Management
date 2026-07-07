package routes

import (
	controller "github.com/Mekde-lawit/Restaurant-Management/src/controllers"
	"github.com/gin-gonic/gin"
)

func OrderRoutes(route *gin.Engine) {
	route.GET("/orders", controller.GetOrders)
	route.GET("/order/:id", controller.GetOrder)
	route.POST("/order", controller.CreateOrder)
	route.PATCH("/order/:id", controller.UpdateOrder)
}
