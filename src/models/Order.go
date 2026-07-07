package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Order struct {
	ID         bson.ObjectID `bson:"_id,omitempty"`
	Order_Id   string        `bson:"order_id" validate:"required"`
	Table_Id   string        `bson:"table_id" validate:"required"`
	Order_Date time.Time     `bson:"order_date"`
	Created_At time.Time     `bson:"created_at"`
	Updated_At time.Time     `bson:"updated_at"`
}
