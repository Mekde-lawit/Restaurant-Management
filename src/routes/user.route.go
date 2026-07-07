package routes

import (
	controller "github.com/Mekde-lawit/Restaurant-Management/src/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(route *gin.Engine) {
	route.GET("/users", controller.GetUsers)
	route.GET("/user/:id", controller.GetUser)
	route.POST("/user/signup", controller.Signup)
	route.POST("/user/login", controller.Login)
}
