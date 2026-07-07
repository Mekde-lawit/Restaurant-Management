package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"time"
)

type Invoice struct {
	ID               bson.ObjectID `bson:"_id,omitempty"`
	Invoice_Id       string        `bson:"invoice_id"`
	Order_Id         string        `bson:"order_id"`
	Payment_Method   *string       `bson:"payment_method" validate:"required,oneof=CARD CASH "`
	Payment_Status   *string       `bson:"payment_status" validate:"required,oneof=PANDING PAID "`
	Payment_Due_Date time.Time     `bson:"payment_due_date"`
	Created_At       time.Time     `bson:"created_at"`
	Updated_At       time.Time     `bson:"updated_at"`
}
