package controllers

import (
	"context"
	"log"
	"math"
	"net/http"
	"time"

	"github.com/Mekde-lawit/Restaurant-Management/src/database"
	"github.com/Mekde-lawit/Restaurant-Management/src/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection("food")
var menuCollection *mongo.Collection = database.OpenCollection("menu")
var validate = validator.New()

func GetFoods(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	cursor, err := foodCollection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	defer cursor.Close(ctx)

	var foods []models.Food

	for cursor.Next(ctx) {
		var food models.Food

		if err := cursor.Decode(&food); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		foods = append(foods, food)
	}

	c.JSON(http.StatusOK, foods)
}

func GetFood(c *gin.Context) {
	ctx, cancel :=context.WithTimeout(context.Background(), 100 * time.Second )
	defer cancel()
	foodId := c.Param("id")
   var food models.Food 

   err := foodCollection.FindOne(ctx, bson.M{"food_id":foodId}).Decode(&food)
   

   if err != nil {
	log.Println("error", err)
	c.JSON(http.StatusNotFound, gin.H{
			"error": "Food not found",
		})
	return
   }
   c.JSON(http.StatusOK,food)
}


func CreateFood(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var food models.Food
	var menu models.Menu

	if err := c.BindJSON(&food); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Incorrect Entry!",
		})
		return
	}

	if err := validate.Struct(food); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Validation failed!",
		})
		return
	}

	err := menuCollection.FindOne(
		ctx,
		bson.M{"menu_id": food.Menu_Id},
	).Decode(&menu)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Menu ID not found!",
		})
		return
	}

	food.Created_At = time.Now()
	food.Updated_At = time.Now()

	food.ID = bson.NewObjectID()
	food.Food_Id = food.ID.Hex()

	num := toFixed(*food.Price, 2)
	food.Price = &num

	_, err = foodCollection.InsertOne(ctx, food)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create food",
		})
		return
	}

	c.JSON(http.StatusCreated, food)
}

func UpdateFood(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	foodId := c.Param("id")

	var updateData bson.M

	if err := c.BindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	// Don't allow these fields to be modified
	delete(updateData, "_id")
	delete(updateData, "food_id")
	delete(updateData, "created_at")

	// Always update the updated_at field
	updateData["updated_at"] = time.Now()

	result, err := foodCollection.UpdateOne(
		ctx,
		bson.M{"food_id": foodId},
		bson.M{
			"$set": updateData,
		},
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Food not found",
		})
		return
	}

	var updatedFood models.Food

	err = foodCollection.FindOne(
		ctx,
		bson.M{"food_id": foodId},
	).Decode(&updatedFood)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedFood)
}



func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return math.Round(num*output) / output
}