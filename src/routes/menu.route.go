package routes

import (
	controller "github.com/Mekde-lawit/Restaurant-Management/src/controllers"
	"github.com/gin-gonic/gin"
)

func MenuRoutes(route *gin.Engine) {
	route.GET("/menus", controller.GetMenus)
	route.GET("/menu/:id", controller.GetMenu)
	route.POST("/menu", controller.CreateMenu)
	route.PATCH("/menu/:id", controller.UpdateMenu)
}
