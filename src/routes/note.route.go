package routes

import (
	controller "github.com/Mekde-lawit/Restaurant-Management/src/controllers"
	"github.com/gin-gonic/gin"
)

func NoteRoutes(route *gin.Engine) {
	route.GET("/notes", controller.GetNotes)
	route.GET("/note/:id", controller.GetNote)
	route.POST("/note", controller.CreateNote)
	route.PATCH("/note/:id", controller.UpdateNote)
}
