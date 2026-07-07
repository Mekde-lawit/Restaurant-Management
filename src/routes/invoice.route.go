package routes

import (
	controller "github.com/Mekde-lawit/Restaurant-Management/src/controllers"
	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(route *gin.Engine) {
	route.GET("/incoices", controller.GetInvoices)
	route.GET("/invoice/:id", controller.GetInvoice)
	route.POST("/invoice", controller.CreateInvoice)
	route.PATCH("/invoice/:id", controller.UpdateInvoice)
}
