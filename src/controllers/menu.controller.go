package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/Mekde-lawit/Restaurant-Management/src/database"
	"github.com/Mekde-lawit/Restaurant-Management/src/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)


var menuCollection *mongo.Collection = database.OpenCollection("menu")

func GetMenus(c *gin.Context) {
	// logic
}

func GetMenu(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	menuId := c.Param("id")

	var menu models.Menu

	err := menuCollection.FindOne(
		ctx,
		bson.M{"menu_id": menuId},
	).Decode(&menu)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Menu not found",
		})
		return
	}

	c.JSON(http.StatusOK, menu)
}

func CreateMenu(c *gin.Context) {
	// logic
}

func UpdateMenu(c *gin.Context) {
	// logic
}
