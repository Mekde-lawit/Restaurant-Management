package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Menu struct {
	ID         bson.ObjectID `bson:"_id,omitempty"`
	Menu_Id    string        `bson:"menu_id" validate:"required"`
	Food_Id    string        `bson:"food_id" validate:"required"`
	Name       string        `bson:"name" validate:"required"`
	Category   string        `bson:"order_id" validate:"required"`
	Start_Date time.Time     `bson:"start_date"`
	End_Date   time.Time     `bson:"end_date"`
	Created_At time.Time     `bson:"created_at"`
	Updated_At time.Time     `bson:"updated_at"`
}
