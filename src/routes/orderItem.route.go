package routes

import (
	controller "github.com/Mekde-lawit/Restaurant-Management/src/controllers"
	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(route *gin.Engine) {
	route.GET("/orderItems", controller.GetOrderItems)
	route.GET("/orderItem/:id", controller.GetOrderItem)
	route.GET("/orderItems-order/:id", controller.GetOrderItemsByOrderId)
	route.POST("/orderItem", controller.CreateOrderItem)
	route.PATCH("/orderItem/:id", controller.UpdateOrderItem)
}
