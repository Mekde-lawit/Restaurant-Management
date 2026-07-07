package main

import (
	"os"

	//db "github.com/Mekde-lawit/Restaurant-Management/src/database"
	env "github.com/Mekde-lawit/Restaurant-Management/src/env"
	route "github.com/Mekde-lawit/Restaurant-Management/src/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	env.InitEnv()
}

func main() {
	port := os.Getenv("PORT")
	router := gin.Default() // default logger and recovery
	route.UserRoutes(router)
	route.MenuRoutes(router)
	route.TableRoutes(router)
	route.OrderRoutes(router)
	route.OrderItemRoutes(router)
	route.FoodRoutes(router)
	route.InvoiceRoutes(router)
	route.FoodRoutes(router)
	route.NoteRoutes(router)

	router.Run(":" + port)
}
