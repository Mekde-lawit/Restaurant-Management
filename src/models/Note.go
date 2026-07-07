package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Note struct {
	ID            bson.ObjectID `bson:"_id,omitempty"`
	Note_Id       string        `bson:"note_id"`
	Title         string       `bson:"title"`
	Text    string      `bson:"text"`
	Created_At    time.Time     `bson:"created_at"`
	Updated_At    time.Time     `bson:"updated_at"`
}
