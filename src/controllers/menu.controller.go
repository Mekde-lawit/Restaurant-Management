package controllers

import (
	"context"
	"log"
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
    ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
    defer cancel()

    cursor, err := menuCollection.Find(ctx, bson.M{})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to fetch menus",
        })
		log.Println("err", err)
        return
    }
    defer cursor.Close(ctx)

    var menus []models.Menu

    if err := cursor.All(ctx, &menus); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to decode menus",
        })
		log.Println("err", err)
        return
    }

    c.JSON(http.StatusOK, menus)
}

func GetMenu(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	id := c.Param("id")

	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid menu ID",
		})
		return
	}

	var menu models.Menu

	err = menuCollection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&menu)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Menu not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch menu",
		})
		return
	}

	c.JSON(http.StatusOK, menu)
}

func CreateMenu(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	var menu models.Menu

	if err := c.ShouldBindJSON(&menu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	now := time.Now().UTC()

	menu.Created_At = now
	menu.Updated_At = now
	menu.ID = bson.NewObjectID()
	menu.Menu_Id = menu.ID.Hex()
	result, err := menuCollection.InsertOne(ctx, menu)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create menu",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Menu created successfully",
		"id":      result.InsertedID,
	})
}


func UpdateMenu(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()
// check id
	id := c.Param("id")
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid menu ID",
		})
		return
	}

	// check body
	var menu models.Menu
	if err := c.ShouldBindJSON(&menu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	update := bson.M{
    "$set": bson.M{
        "name":       menu.Name,
        "category":   menu.Category,
        "food_id":    menu.Food_Id,
        "start_date": menu.Start_Date,
        "end_date":   menu.End_Date,
        "updated_at": time.Now().UTC(),
    },
}

	result, err := menuCollection.UpdateOne(
		ctx,
		bson.M{"_id": objectID},
		update,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update menu",
		})
		return
	}

	if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Menu not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Menu updated successfully",
	})
}