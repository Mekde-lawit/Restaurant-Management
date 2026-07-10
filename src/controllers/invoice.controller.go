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

var invoiceCollection = database.OpenCollection("invoice")

// GET /invoices
func GetInvoices(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	cursor, err := invoiceCollection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch invoices",
		})
		return
	}
	defer cursor.Close(ctx)

	var invoices []models.Invoice

	if err := cursor.All(ctx, &invoices); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to decode invoices",
		})
		return
	}

	c.JSON(http.StatusOK, invoices)
}

// GET /invoices/:id
func GetInvoice(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	id := c.Param("id")

	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid invoice ID",
		})
		return
	}

	var invoice models.Invoice

	err = invoiceCollection.FindOne(
		ctx,
		bson.M{"_id": objectID},
	).Decode(&invoice)

	if err != nil {

		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Invoice not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch invoice",
		})
		return
	}

	c.JSON(http.StatusOK, invoice)
}

// POST /invoices
func CreateInvoice(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	var invoice models.Invoice

	if err := c.ShouldBindJSON(&invoice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	now := time.Now().UTC()

	invoice.ID = bson.NewObjectID()
	invoice.Invoice_Id = invoice.ID.Hex()
	invoice.Created_At = now
	invoice.Updated_At = now

	result, err := invoiceCollection.InsertOne(ctx, invoice)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create invoice",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Invoice created successfully",
		"id":      result.InsertedID,
	})
}

// PUT /invoices/:id
func UpdateInvoice(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	id := c.Param("id")

	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid invoice ID",
		})
		return
	}

	var invoice models.Invoice

	if err := c.ShouldBindJSON(&invoice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	update := bson.M{
		"$set": bson.M{
			"order_id":          invoice.Order_Id,
			"payment_method":    invoice.Payment_Method,
			"payment_status":    invoice.Payment_Status,
			"payment_due_date":  invoice.Payment_Due_Date,
			"updated_at":        time.Now().UTC(),
		},
	}

	result, err := invoiceCollection.UpdateOne(
		ctx,
		bson.M{"_id": objectID},
		update,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update invoice",
		})
		return
	}

	if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Invoice not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Invoice updated successfully",
	})
}