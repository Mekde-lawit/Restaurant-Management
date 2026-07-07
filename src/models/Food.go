package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"time"
)

type Food struct {
	ID         bson.ObjectID `bson:"_id,omitempty"`
	Name       *string       `bson:"name" validate:"required,min=2,max=100"`
	Price      *float64      `bson:"price" validate:"required"`
	Food_Image *string       `bson:"food_image" validate:"required"`
	Food_Id    string        `bson:"food_id"`
	Menu_Id    *string       `bson:"menu_id"`
	Created_At time.Time     `bson:"created_at"`
	Updated_At time.Time     `bson:"updated_at"`
}
