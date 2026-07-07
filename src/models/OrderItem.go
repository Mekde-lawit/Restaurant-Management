package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type OrderItem struct {
	ID            bson.ObjectID `bson:"_id,omitempty"`
	Order_Item_Id string        `bson:"order_item_id"`
	Size          *string       `bson:"size" validate:"required,oneof=S M L"`
	Unit_Price    *float64      `bson:"unit_price" validate:"required"`
	Food_Id       *string       `bson:"food_id" validate:"required"`
	Order_Id      string        `bson:"order_id" validate:"required"`
	Created_At    time.Time     `bson:"created_at"`
	Updated_At    time.Time     `bson:"updated_at"`
}
