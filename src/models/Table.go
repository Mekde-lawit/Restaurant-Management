package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Table struct {
	ID              bson.ObjectID `bson:"_id,omitempty"`
	Number_Of_Guest *int          `bson:"no_of_guest" validate:"required"`
	Table_Number    *int          `bson:"table_no" validate:"required"`
	Table_Id        string        `bson:"table_id" validate:"required"`
	Created_At      time.Time     `bson:"created_at"`
	Updated_At      time.Time     `bson:"updated_at"`
}
